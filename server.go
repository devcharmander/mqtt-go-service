package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	emitter "github.com/emitter-io/go/v2"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	mqtt   *mqttClient
	router *httprouter.Router
}

type mqttClient struct {
	client *emitter.Client
	host   string
	topic  string
	key    string
}

func NewServer() *server {
	mc := &mqttClient{
		host:  "tcp://127.0.0.1:8081",
		topic: "semicolons-demo/",
		key:   "P3W_EzP8IMQsqSF5Hw1W3OUn0aH44zPP",
	}

	client, err := emitter.Connect(mc.host, onConnect)
	if err != nil {
		log.Fatalf("Connection failed %v", err)
	}
	mc.client = client

	return &server{
		mqtt:   mc,
		router: httprouter.New(),
	}
}

func (srv *server) SetHandlers() {
	srv.mqtt.client.OnPresence(onPresence)
}

func (srv *server) httpPublishHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body. Error %v", err)
		fmt.Fprint(w, false)
	}
	msg := fmt.Sprintf("%s", body)

	if err := srv.mqtt.client.Publish(srv.mqtt.key, srv.mqtt.topic, msg); err != nil {
		log.Printf("Error publishing the message. Error %v", err)
		fmt.Fprint(w, false)
	}
	fmt.Fprint(w, true)
}
