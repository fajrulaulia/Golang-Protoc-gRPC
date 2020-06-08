package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

func (api *API) getByName(title string, client *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}
	*client = getItem
	return nil
}

func (api *API) CreateItem(item Item, client *Item) error {
	database = append(database, item)
	log.Println("CREATE ITEM OK")
	*client = item
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Println("rpc.Register.error", err)
		return
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println("net.Listen.error", err)
		return
	}
	log.Println("RPC Server run", 6969)

	err = http.Serve(listen, nil)
	if err != nil {
		log.Println("http.Serve.error", err)
		return
	}
}
