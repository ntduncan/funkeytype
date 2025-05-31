package session

type Session struct {
	sessionBest string
}

func (s *Session) SetSessionBest(newVal string) {
	s.sessionBest = newVal
}

func (s *Session) GetSessionBest() string {
	return s.sessionBest
}

