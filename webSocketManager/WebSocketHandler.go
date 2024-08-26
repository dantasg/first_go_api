package webSocketManager

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

// Broadcast envia uma mensagem para todos os clientes conectados.
type MessageReceived struct {
	Event string
	Data  interface{}
}

// WebSocketHandler lida com as conexões WebSocket.
func WebSocketHandler(cm *ClientManager) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		defer c.Close()

		log.Println(c)

		clientID := c.Query("id")
		if clientID == "" {
			log.Println("ID do cliente não fornecido")
			return
		}

		cm.AddClient(clientID, c)
		log.Println("Cliente %s conectado", clientID)

		defer cm.RemoveClient(clientID)

		eventMessage := Message{
			Event: "connection",
			Data:  "UserId = " + clientID + ", connected success!",
		}

		messageBytes, err := json.Marshal(eventMessage)
		if err != nil {
			log.Println("Erro ao serializar a mensagem:", err)
		}

		cm.SendToClient(clientID, []byte(messageBytes))

		// Mantém a conexão aberta para receber mensagens do cliente
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Erro de leitura:", err)
				break
			}

			var receivedMessage MessageReceived
			if err := json.Unmarshal(msg, &receivedMessage); err != nil {
				log.Println("Erro ao desserializar mensagem:", err)
				continue
			}

			log.Printf("Mensagem recebida de %s - Evento: %s, Dados: %s", clientID, receivedMessage.Event, receivedMessage.Data)

			// Lógica para lidar com diferentes tipos de eventos
			switch receivedMessage.Event {
			case "hello":
				log.Println("Tratando evento someEvent com dados:", receivedMessage.Data)

				// var loginUser LoginUser
				// dataBytes, err := json.Marshal(receivedMessage.Data)
				// if err != nil {
				// 	log.Println("Erro ao serializar Data:", err)
				// 	continue
				// }

				// if err := json.Unmarshal(dataBytes, &loginUser); err != nil {
				// 	log.Println("Erro ao desserializar Data:", err)
				// 	continue
				// }

				// log.Printf("Login recebido - Email: %s, Password: %s", loginUser.Email, loginUser.Password)

				break

			case "anotherEvent":
				log.Println("Tratando evento anotherEvent com dados:", receivedMessage.Data)
				break

			default:
				log.Println("Evento não reconhecido:", receivedMessage.Event)
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("Erro ao enviar mensagem:", err)
				break
			}
		}
	}
}
