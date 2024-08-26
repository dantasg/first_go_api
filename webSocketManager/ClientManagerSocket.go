package webSocketManager

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type ClientManager struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

// Broadcast envia uma mensagem para todos os clientes conectados.
type Message struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*websocket.Conn),
	}
}

func (cm *ClientManager) AddClient(id string, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[id] = conn
}

func (cm *ClientManager) RemoveClient(id string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, id)
}

func (cm *ClientManager) Broadcast(message []byte) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Erro ao desserializar mensagem:", err)
		return
	}

	fmt.Printf("Broadcast - Evento: %s, Dados: %s\n", msg.Event, msg.Data)

	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, conn := range cm.clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Erro ao enviar mensagem:", err)
		}
	}
}

// SendToClient envia uma mensagem para um cliente específico.
func (cm *ClientManager) SendToClient(id string, message []byte) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Erro ao desserializar mensagem:", err)
		return
	}

	fmt.Printf("Broadcast - Evento: %s, Dados: %s\n", msg.Event, msg.Data)

	cm.mu.Lock()
	defer cm.mu.Unlock()
	if conn, ok := cm.clients[id]; ok {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Erro ao enviar mensagem:", err)
		}
	}
	return
}
