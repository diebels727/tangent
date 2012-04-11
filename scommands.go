package tangent

import (
	"github.com/xales/tangent/irc"
	"strings"
)

type scommand struct {
	args   int
	handle func(*irc.Line, *Server)
}

var scommands = map[string]*scommand{
	"PING": {0, s_ping},
	"001":  {1, s_001},
	"433":  {2, s_433},
	"CAP":  {2, s_cap},
	"JOIN": {1, s_join},
	"PART": {1, s_part},
}

var caps = map[string]bool{
	"account-notify": true,
	"extended-join":  true,
	"multi-prefix":   true,
}

func s_cap(l *irc.Line, s *Server) {
	switch l.Args[1] {
	case "LS":
		var encap []string
		for _, c := range strings.Fields(l.Args[2]) {
			if caps[c] {
				encap = append(encap, c)
			}
		}
		if len(encap) > 0 {
			s.Cmd("CAP", "REQ", strings.Join(encap, " "))
			s.RawCmd("CAP", "END")
			s.register()
		}
	case "ACK":
		for _, c := range strings.Fields(l.Args[2]) {
			s.capabs[c] = true
		}
	case "NAK":
		for _, c := range strings.Fields(l.Args[2]) {
			s.capabs[c] = false
		}
	}
}

func s_join(l *irc.Line, s *Server) {
	u := s.fuser(&l.HostMask)
	if s.capabs["extended-join"] && len(l.Args) > 2 {
		u.Account, u.Real = l.Args[1], l.Args[2]
	}
	ch := s.fchan(l.Args[0])
	u.Channels[ch] = true
	ch.Users[u] = true
}

func s_part(l *irc.Line, s *Server) {
	u := s.fuser(&l.HostMask)
	ch := s.fchan(l.Args[0])
	delete(u.Channels, ch)
	delete(ch.Users, u)
}

func s_ping(l *irc.Line, s *Server) {
	s.Cmd("PONG", l.Args...)
}

func s_001(l *irc.Line, s *Server) {
	s.us.Nick = l.Args[0]
	s.Name = l.Src
	s.context.conns++ // FIXME This should probably be elsewhere
	s.Users[strings.ToLower(s.us.Nick)] = s.us
	s.e(new(Connect))
}

func s_433(l *irc.Line, s *Server) {
	newnick := l.Args[1] + "_"
	logger.Infof("%s already in use, trying %s", l.Args[1], newnick)
	s.Nick(newnick)
}
