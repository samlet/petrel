package dummy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/samlet/petrel/alfin/cache"
	"time"
)

type ValueProcs struct{
	client *redis.Client
}

// SetData
// @serial 1
// @message "Cannot execute task {{.Name}}"
// @when:
// 	key LIKE "data_*"
func (t ValueProcs) SetData(key string, valueList interface{}) error {
	valueList, err := json.Marshal(valueList)
	if err != nil {
		return fmt.Errorf("cannot marshal value: %w", err)
	}
	err = t.client.Set(key, valueList, 0).Err()
	if err != nil {
		return fmt.Errorf("cannot store value: %w", err)
	}
	return nil
}

// SetValue
// @serial 1
// @message "Cannot execute task {{.Name}}"
// @when:
// 	key LIKE "state_*"
func SetValue(ctx context.Context, key string, value string) {
	store:=cache.FromContext(ctx)
	store.Put(key, value)
}

// GetValue get a key's value
// @timeout 12s
// @serial 2
func GetValue(ctx context.Context, key string) string {
	store:=cache.FromContext(ctx)
	return store.Get(key).(string)
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
// @serial 3
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

