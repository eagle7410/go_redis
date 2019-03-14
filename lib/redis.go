package lib

import (
	"fmt"
	"github.com/go-redis/redis"
)

type (
	redisDb struct {
		Client *redis.Client
	}
)

var Redis redisDb

func (i *redisDb) Init () error {

	i.Client = redis.NewClient(&redis.Options{
		Addr:     ENV.HOST + ":" + ENV.PORT, // use default Addr
		Password: ENV.AUTH_PASS,               // no password set
		DB:       0,                // use default DB
	})

	answer, err := i.Client.Ping().Result()

	fmt.Printf("Be PING, Answer %v \n", answer )

	return err
}
