package dummy

import (
	"context"
	"github.com/samlet/petrel/alfin/cache"
	"time"
)

// GetValue get a key's value
// @timeout 12s
// @serial 2
func GetValue(ctx context.Context, key string) string {
	store:=cache.FromContext(ctx)
	return store.Get(key).(string)
}

// SetValue
// @serial 2
// @message "Cannot execute task {{.Name}}"
// @when:
// 	key LIKE "state_*"
func SetValue(ctx context.Context, key string, value string) {
	store:=cache.FromContext(ctx)
	store.Put(key, value)
}

// SetPrice
// @serial 1.1
// @message "Cannot execute task {{.Name}}"
// @when:
// 	key LIKE "price_*"
// 	price BETWEEN 1 AND 1000
func SetPrice(ctx context.Context, key string, price int64) {
	store:=cache.FromContext(ctx)
	store.Put(key, price)
	store.Put(key+"_ts", time.Now())
}

// GetPrice
// @serial 1.1
// @message "Cannot execute task {{.Name}}"
// @when:
// 	key LIKE "price_*"
// @validate:
// 	AND (
//		result._1 > 0
// 		result._2 < now()
// 	)
func GetPrice(ctx context.Context, key string) (int64, time.Time){
	store:=cache.FromContext(ctx)
	return store.Get(key).(int64), store.Get(key+"_ts").(time.Time)
}

