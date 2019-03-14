package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_redis/lib"
	"log"
	"time"
)

func init() {
	if err:= lib.ENV.Init(); err != nil {
		log.Fatal(err)
	}

	if err:= lib.Redis.Init(); err != nil {
		log.Fatal(err)
	}

}
func main() {
	fmt.Println("RUN")

	if err := lib.Redis.Client.Set("Test", 12, 0).Err(); err != nil {
		log.Fatal(err)
	}

	val, err := lib.Redis.Client.Get("Test").Result()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Val is %t \n", val)

	val, err = lib.Redis.Client.Get("Test1").Result()

	if err == redis.Nil {
		fmt.Printf("Test1 not exist \n")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Val is %t \n", val)
	}

	pubsub := lib.Redis.Client.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.

	if _, err := pubsub.Receive(); err != nil {
		log.Fatal(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.

	if err = lib.Redis.Client.Publish("mychannel1", "hello").Err(); err != nil {
		log.Fatal(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
