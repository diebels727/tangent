package tangent

import (
	"github.com/xales/tangent/irc"
)

type scommand struct {
	args   int
	handle func(*irc.Line, *Server)
}

var scommands = map[string]*scommand{
	"PING": {0, s_ping},
    "001": {1, s_001},
    "433": {1, s_433},
}

func s_ping(l *irc.Line, s *Server) {
	s.Cmd("PONG", l.Args...)
}

func s_001(l *irc.Line, s *Server) {
    s.connected = true
    s.us.Nick = l.Args[0]
}

func s_433(l *irc.Line, s *Server) {
    newnick := l.Args[1] + "_"
    logger.Infof("%s already in use, trying %s", l.Args[1], newnick)
    s.Nick(newnick)
}
