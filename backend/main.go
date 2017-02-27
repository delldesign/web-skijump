package main

import (
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/websocket"

	log "github.com/Sirupsen/logrus"
)

type message struct {
	// the json tag means this will serialize as a lowercased field
	Message string `json:"message"`
}

type object struct {
	value int
}

func (o *object) socket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message

		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}

		log.Println("Received message:", m.Message)
		o.value++
		// send a response
		m2 := message{strconv.Itoa(o.value)}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")

	var obj = object{0}

	http.Handle("/ws", websocket.Handler(obj.socket))

	type message struct {
		// the json tag means this will serialize as a lowercased field
		Message string `json:"message"`
	}
	http.ListenAndServe(":9000", nil)
}
