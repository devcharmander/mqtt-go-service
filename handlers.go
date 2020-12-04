package main

import (
	"fmt"

	emitter "github.com/emitter-io/go/v2"
)

func onPresence(client *emitter.Client, ev emitter.PresenceEvent) {
	fmt.Println("Presence event:", len(ev.Who), "subscriber(s) at channel", ev.Channel)
}

func onConnect(client *emitter.Client, msg emitter.Message) {
	fmt.Printf("Recieved message %s from topic '%s'\n", msg.Payload(), msg.Topic())
}

func onSubscribe(client *emitter.Client, msg emitter.Message) {
	fmt.Printf("Recieved message on subscribe handler '%s' from topic '%s'\n", msg.Payload(), msg.Topic())
}
