package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type clientPool struct {
	id      uuid.UUID
	clients map[*client]bool
}

type client struct {
	conn        *websocket.Conn
	send        chan []string
	client_pool *clientPool
}

func make_pool() (*clientPool, uuid.UUID) {

	uuid := uuid.New()

	client_pool := clientPool{
		id:      uuid,
		clients: make(map[*client]bool),
	}

	return &client_pool, uuid
}

func (cp *clientPool) add_to_pool(client *client) {
	cp.clients[client] = true
}

func (cp *clientPool) get_total_pool_members() int {
	return len(cp.clients)
}
