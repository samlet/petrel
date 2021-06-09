package roles


import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	uuid "github.com/satori/go.uuid"
	"log"
	"reflect"
	"sync"

	"github.com/go-redis/redis"
)
// Watcher
type Watcher struct {
	options options // 选项
	pubConn  redis.UniversalClient //生产消息
	subConn  redis.UniversalClient //订阅消息
	callback func(string) // 回调方法
	closed   chan struct{}
	once     sync.Once
}
// 选项
type options struct {
	id string // 当前节点的id
	channel string // redis的管道名称
}
// Option 选项接口
type Option interface {
	Set(options *options)
}
// channel选项
type channelOption string
func (o channelOption) Set(opts *options)  {
	opts.channel = string(o)
}
func WithChannel(channel string) Option {
	return channelOption(channel)
}
// id选项
type idOption string
func (o idOption) Set(opts *options)  {
	opts.id = string(o)
}
func WithID(id string) Option {
	return idOption(id)
}
// NewWatcher 实例化
func NewWatcher(conn redis.UniversalClient, opts ...Option) (persist.Watcher, error) {
	w := &Watcher{
		closed: make(chan struct{}),
		pubConn: conn,
		subConn: conn,
	}

	w.options = options{
		channel:  "/casbin",
		id : uuid.NewV4().String(),
	}

	for _, opt := range opts {
		opt.Set(&w.options)
	}

	go func() {
		for {
			select {
			case <-w.closed:
				return
			default:
				err := w.subscribe()
				if err != nil {
					fmt.Printf("Failure from Redis subscription: %v", err)
				}
			}
		}
	}()

	return w, nil
}

// SetUpdateCallBack sets the update callback function invoked by the watcher
// when the policy is updated. Defaults to Enforcer.LoadPolicy()
func (w *Watcher) SetUpdateCallback(callback func(string)) error {
	w.callback = callback
	return nil
}

// Update publishes a message to all other casbin instances telling them to
// invoke their update callback
func (w *Watcher) Update() error {
	// 当有权限更新时，发送当前节点的ID,通知其他节点有更新
	if _, err := w.pubConn.Publish(w.options.channel, w.options.id).Result(); err != nil {
		return err
	}
	return nil
}

// Close disconnects the watcher from redis
func (w *Watcher) Close() {
	w.once.Do(func() {
		close(w.closed)
		_ = w.subConn.Close()
		_ = w.pubConn.Close()
	})
}
// subscribe 订阅redis
func (w *Watcher) subscribe() error {
	psc := w.subConn.Subscribe(w.options.channel)
	defer psc.Unsubscribe()
	for {
		receive, err := psc.Receive()
		fmt.Println(reflect.TypeOf(receive))
		if err != nil {
			fmt.Printf("try subscribe channel[%s] error[%s]\n", w.options.channel, err.Error())
			return nil
		}
		switch n := receive.(type) {
		case error:
			return n
		case *redis.Message:
			// 需要跳过当前生产者
			if n.Payload !=  w.options.id && w.callback != nil {
				w.callback("success") // 执行回调
			}

		case *redis.Subscription:
			if n.Count == 0 {
				return nil
			}
		}
	}
}

// ?
// ref: https://hongker.github.io/2021/02/19/golang-rbac-watcher/
func StartWathcer(){
	authEnforcer, err := casbin.NewEnforcer("rbac/conf/rbac_model.conf")
	if err != nil {
		log.Fatal(err)
	}

	redisConn := redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1",
			Password: "",
			DB:       0,
		})

	if _, err := redisConn.Ping().Result(); err != nil {
		log.Fatal(err)
	}

	watcher, err := NewWatcher(redisConn)
	if err != nil {
		log.Fatal(err)
	}
	watcher.SetUpdateCallback(func(s string) {
		// 当有权限更新时，会执行这个回调，具体更新内存的权限逻辑可以自定义
		authEnforcer.LoadPolicy()
	})
	authEnforcer.SetWatcher(watcher)
}