package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Hub_list struct {
	hub map[uuid.UUID]*clientPool
}

var pool_list = Hub_list{
	hub: make(map[uuid.UUID]*clientPool),
}

func main() {

	http.HandleFunc("/", hello_world_handle)

	//Dummy Default Room
	pool_id := uuid.New()
	dummy_pool := clientPool{
		id:      pool_id,
		clients: make(map[*client]bool),
	}

	pool_list.hub[pool_id] = &dummy_pool

	http.HandleFunc("/sendpools", func(w http.ResponseWriter, r *http.Request) {
		send_pools(w, r)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket_handle(w, r)
	})

	http.HandleFunc("/addhub", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		new_uid := uuid.New()
		pool_list.hub[new_uid] = &clientPool{
			id:      new_uid,
			clients: make(map[*client]bool),
		}
		w.Write([]byte(new_uid.String()))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
