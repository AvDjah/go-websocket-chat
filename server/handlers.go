package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type client_message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type response_message struct {
	Type string `json:"type"`
	Data string `json:"data"`
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
		var resp response_message

		if pool_id_req.Type == "MESSAGE" {
			resp.Type = "MESSAGE"
			resp.Data = "No Pool Joined"
			if client.client_pool == nil {
				conn.WriteJSON(resp)
				return
			}

			fmt.Println("Message to Send: " + pool_id_req.Data)
			client.broadcast_message(pool_id_req.Data)
		} 

		if(pool_id_req.Type == "POOL_ID_JOIN"){
			fmt.Println("Request to Join Pool")
			err := client.join_pool(pool_id_req.Data)

			resp.Type = "POOL_ID_JOIN_RESULT"
			if err != nil {
				resp.Data = "-1"
			} else {
				uid,err := uuid.Parse(pool_id_req.Data)
				if err != nil {
					log.Fatal("UUID parse Error")
				}
				resp.Data = pool_list.hub[uid].name
			}
			err = conn.WriteJSON(resp)
			// msg,err := json.Marshal(resp)
			if err != nil {
				log.Fatal("Marshalling Error")
			}
		}


		// if pool_id_req.Pool_id != "" {
		// 	client.join_pool(pool_id_req.Pool_id)
		// } else {
		// 	if client.client_pool == nil {
		// 		conn.WriteMessage(1, []byte("No Pool Joined"))
		// 	} else {
		// 		fmt.Println("Message to Send: " + pool_id_req.Message)
		// 		client.broadcast_message(pool_id_req.Message)
		// 	}
		// }

		// fmt.Print(pool_id_req.Pool_id)
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

	var resp response_message
	resp.Type = "MESSAGE"
	resp.Data = "SEND_POOL_UUID"
	conn.WriteJSON(resp)
	go handle_websocket_conn(conn, &client)
}


type pool_resp struct {
	Name  string `json:"name"`
	Pool_id string `json:"pool_id"`
}


func send_pools(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var pool_uuid_list []pool_resp
	for pool_id := range pool_list.hub {
		pool_uuid_list = append(pool_uuid_list, pool_resp{
			Name: pool_list.hub[pool_id].name,
			Pool_id: pool_id.String(),
		})
	}

	json_list, err := json.Marshal(pool_uuid_list)

	if err != nil {
		log.Fatal("Error Marshalling")
	}
	fmt.Println(pool_uuid_list)
	w.Write(json_list)
}


func add_pool(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Add Pool Err: 1",err)
	}

	var msg client_message
	err = json.Unmarshal(body,&msg)
	if err != nil {
		log.Fatal(err)
	}

	uuid := uuid.New()
	fmt.Println("resp : ",msg.Data)
	pool_list.hub[uuid] = &clientPool{
		name: msg.Data,
		id: uuid,
		clients: make(map[*client]bool),
	}

}