package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

//LLqAKaY1jHtdaoKXoORfIdKv-j7DIryW
//const key = "TWyNGb5sAODJnWKa8kJO1caLiVNbuZsc"

var isPublish bool

func main() {
	//	done := make(chan bool)
	flag.BoolVar(&isPublish, "p", false, "Set this flag to true to run a publisher")
	flag.Parse()
	server := NewServer()
	server.SetHandlers()

	server.router.POST("/mqtt", server.httpPublishHandler)

	mqtt := server.mqtt

	fmt.Println("Client name:", mqtt.client.ID())

	fmt.Println("Asking for presence")
	if err := mqtt.client.Presence(mqtt.key, mqtt.topic, true, false); err != nil {
		log.Fatalf("Could not determine the presence. Error %v", err)
	}
	if isPublish {
		for {
			msg := ""
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter your message:")
			msg, _ = reader.ReadString('\n')
			msg = strings.Replace(msg, "\n", "", 1)
			fmt.Println("msg:", msg)
			mqtt.client.Publish(mqtt.key, mqtt.topic, msg)
		}
	}
	fmt.Printf("Subscring to %s\n", mqtt.topic)
	mqtt.client.Subscribe(mqtt.key, mqtt.topic, onSubscribe)
	log.Fatal(http.ListenAndServe(":8080", server.router))
	//	<-done
}
