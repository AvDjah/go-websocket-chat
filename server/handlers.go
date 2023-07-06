package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type client_message struct {
	Message string `json:"message"`
	Pool_id string `json:"pool_id"`
}

var upgrader = websocket.Upgrader{}

func hello_world_handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Heelo")
}

func handle_websocket_conn(conn *websocket.Conn, client *client) {
	defer conn.Close()
	defer client.remove_from_pool()
	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		fmt.Println(len(pool_list.hub))

		var pool_id_req client_message
		json.Unmarshal(message, &pool_id_req)

		if pool_id_req.Pool_id != "" {
			client.join_pool(pool_id_req.Pool_id)
		} else {
			if client.client_pool == nil {
				conn.WriteMessage(1, []byte("No Pool Joined"))
			} else {
				fmt.Println("Message to Send: " + pool_id_req.Message)
				client.broadcast_message(pool_id_req.Message)
			}
		}

		fmt.Print(pool_id_req.Pool_id)
	}
}

func websocket_handle(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.RemoteAddr)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := client{
		conn:        conn,
		send:        make(chan []string),
		client_pool: nil,
	}

	conn.WriteMessage(1, []byte("SEND_POOL_UUID"))
	go handle_websocket_conn(conn, &client)
}

func send_pools(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var pool_uuid_list []string
	for pool_id := range pool_list.hub {
		pool_uuid_list = append(pool_uuid_list, pool_id.String())
	}
	json_list, err := json.Marshal(pool_uuid_list)

	if err != nil {
		log.Fatal("Error Marshalling")
	}
	fmt.Println(pool_uuid_list)
	w.Write(json_list)
}
