package main

import (
	"io"
	"log"
	"net/http"

	"github.com/twilio/twilio-go/twiml"
)

func (server *Server) handlerConvo(w http.ResponseWriter, r *http.Request) {
	msg := &twiml.MessagingMessage{}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	incomingMsg := r.PostForm.Get("Body")
	client := r.PostForm.Get("From")
	clientID := generateHash(client)

	server.mut.Lock()
	session, exists := server.sessions[clientID]
	if !exists {
		session = &SessionTracker{
			ClientID:     clientID,
			ClientNumber: client,
		}

		server.sessions[clientID] = session
	}
	server.mut.Unlock()

	if incomingMsg == "Hi" || incomingMsg == "hi" {
		msg.Body = "Hi!"
		session.ClientMSGS = append(session.ClientMSGS, incomingMsg)
	} else if incomingMsg == "goodbye" || incomingMsg == "Goodbye" {
		msg.Body = "See you later nerd"
		session.ClientMSGS = append(session.ClientMSGS, incomingMsg)
		server.endSession(clientID)
	} else {
		msg.Body = "Tu quien eres pues"
		session.ClientMSGS = append(session.ClientMSGS, incomingMsg)
	}

	twiml, err := twiml.Messages([]twiml.Element{msg})
	if err != nil {
		log.Fatal(err)
	}

	server.mut.Lock()
	_, exists = server.sessions[clientID]
	if !exists {
		log.Print("ya se borro viejon, no hay nada ya")
		log.Print(session)
		log.Print(server.sessions[clientID])
	} else {
		log.Print(*session)
	}
	server.mut.Unlock()

	w.Header().Add("Content-Type", "text/xml")
	w.WriteHeader(200)
	io.WriteString(w, twiml)
}
