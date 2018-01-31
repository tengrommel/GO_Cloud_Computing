package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
)

func ExampleNewClient()  {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password:"",
		DB:0,
	})

	pong, err := client.Ping().Result()
	if err !=nil{
		log.Fatal("Redis can't connect")
	}
	fmt.Println(pong)

	err = client.Set("key", "value2", 0).Err()
	if err !=nil{
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err !=nil{
		panic(err)
	}
	fmt.Println(val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil{
		fmt.Println("key2 does not exits")
	} else if err !=nil{
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func main() {
	ExampleNewClient()

}
