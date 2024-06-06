package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/twilio/twilio-go/twiml"
)

func (server *Server) handlerConvo(w http.ResponseWriter, r *http.Request) {
	msg := &twiml.MessagingMessage{}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	incomingMsg := strings.ToLower(r.PostForm.Get("Body"))
	client := r.PostForm.Get("From")
	clientID := generateHash(client)

	server.mut.Lock()
	session, exists := server.sessions[clientID]
	server.mut.Unlock()

	session = server.handleIncommingSession(session, exists, client, clientID)
	session.ClientsPrevMSG = incomingMsg

	if incomingMsg == "hi" {
		msg.Body = "Hi!"
		session.ClientMSGS = append(session.ClientMSGS, incomingMsg)
	} else if incomingMsg == "goodbye" {
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
	session, exists = server.sessions[clientID]
	server.mut.Unlock()

	if !exists {
		log.Print("ya se borro viejon, no hay nada ya")
	} else {
		if session == nil {
			log.Fatal("nil pointer my guy")
		} else {
			jsonData, err := json.MarshalIndent(*session, "", "    ")
			if err != nil {
				log.Fatal("Error marshaling to JSON:", err)
			}
			log.Print(string(jsonData))
		}
	}

	w.Header().Add("Content-Type", "text/xml")
	w.WriteHeader(200)
	io.WriteString(w, twiml)
}

func (server *Server) handleIncommingSession(session *SessionTracker, exists bool, client string, clientID string) *SessionTracker {
	if !exists {
		session = &SessionTracker{
			ClientID:     clientID,
			ClientNumber: client,
		}

		server.mut.Lock()
		if existingSession, exists := server.sessions[clientID]; exists {
			session = existingSession
		} else {
			server.sessions[clientID] = session
		}
		server.mut.Unlock()
	}
	return session
}

func (server *Server) endSession(clientID string) {
	server.mut.Lock()
	delete(server.sessions, clientID)
	server.mut.Unlock()
}
