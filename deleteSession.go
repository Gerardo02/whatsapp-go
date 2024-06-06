package main

func (s *Server) endSession(clientID string) {
	s.mut.Lock()
	delete(s.sessions, clientID)
	s.mut.Unlock()
}
