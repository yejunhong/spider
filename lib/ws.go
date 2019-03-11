package lib

import (
	"time"
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"github.com/gorilla/websocket"
)

type Client struct{
	Id string
	Socket *websocket.Conn // 客户端 连接
}

var upGrader = websocket.Upgrader{  
	CheckOrigin: func (r *http.Request) bool {  
	   return true  
	},  
  }

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	u1 = uuid.NewV4().String()

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}