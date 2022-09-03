package nosql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestExampleClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":5555",
		Password: "", // no password set
		Username: "",
		DB:       0, // use default DB
	})

	err := rdb.Set(ctx, "key", time.Now().String(), 30*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
	pubsub := rdb.Subscribe(ctx, "mychannel")
	defer pubsub.Close()

	ch := pubsub.Channel()
	go func() {
		for i := 0; i < 10; i++ {
			msg := time.Now().String()
			err = rdb.Publish(ctx, "mychannel", msg).Err()
			fmt.Println("mychannel Publish ", msg, err)
			time.Sleep(1 * time.Second)
		}
	}()

	for v := range ch {
		fmt.Println("mychannel body", v.Channel, v.Pattern, v.Payload, v.PayloadSlice)
	}
}
