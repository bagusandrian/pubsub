package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Everything is fine!!!")
	}
	defer c.Close()

	// This is for Publisher
	/*
		Commands:
		redis-cli
		PUBLISH example test
	*/
	// c.Do("PUBLISH", "example", "hello "+time.Now().String())
	/// End here

	// This is for Publisher
	psc := redis.PubSubConn{Conn: c}

	// set subcribe to channel name, on this case is 'example'.
	/*
		Commands:
		redis-cli
		SUBCRIBE example
	*/
	psc.Subscribe("example")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			message := string(v.Data)
			fmt.Println(message)
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
	/// End here

}
