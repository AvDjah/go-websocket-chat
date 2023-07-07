package main

import (
	"fmt"
	"log"
	// "strconv"

	"github.com/google/uuid"
)

func (c *client) broadcast_message(message string) {

	// pool_size := len(c.client_pool.clients)
	// fmt.Println("Pool Size" + )
	// message = message + "->" + strconv.Itoa(pool_size)
	var resp response_message
	resp.Type = "MESSAGE3"
	resp.Data = message
	fmt.Println("Sending: " + message)
	for client := range c.client_pool.clients {
		if c == client {
			continue
		}
		client.conn.WriteJSON(resp)
	}
}

func (c *client) join_pool(id string)  error {
	fmt.Println("joinpool: " + id)
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Error converting string to UUID")
		log.Fatal(err)
		return err
	}
	pool := pool_list.hub[uid]
	pool.clients[c] = true
	c.client_pool = pool
	return nil
}

func (c *client) remove_from_pool() {

	if c.client_pool != nil {
		id := c.client_pool.id
		pool := pool_list.hub[id]
		delete(pool.clients, c)
	}

}
