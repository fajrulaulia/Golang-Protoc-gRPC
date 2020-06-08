package main

import (
	"log"
	"net/rpc"
)

type item struct {
	Title string
	Body  string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	a := item{"First", "body 11"}
	b := item{"Second", "body 22"}
	c := item{"Third", "body 33"}

	var sendFeed item
	client.Call("API.CreateItem", a, &sendFeed)
	client.Call("API.CreateItem", b, &sendFeed)
	client.Call("API.CreateItem", c, &sendFeed)

	client.Call("API.getByName", "First", &sendFeed)
	log.Println("GetByName", sendFeed)

}
