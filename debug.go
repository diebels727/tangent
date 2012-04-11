package tangent

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

// Debug allows controlling using a proxy
type Debug struct {
	lis net.Listener
	cl  []net.Conn
	s   *Server
}

// Debug starts debugging for this server
func (s *Server) Debug(addr string) (dbg *Debug, err error) {
	nc, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	dbg = &Debug{lis: nc, s: s}
	logger.Info("Debugger listening on ", addr)
	go dbg.listen()
	return
}

func (d *Debug) listen() {
	for {
		cl, err := d.lis.Accept()
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Debug("New debug client from ", cl.RemoteAddr())
		go d.serve(cl)
	}
}

func (d *Debug) serve(c net.Conn) {
	br := bufio.NewReader(c)
	for {
		line, prefix, err := br.ReadLine()
		if err != nil {
			logger.Error("DebugClient: ", err)
			return
		}
		if prefix {
			continue
		}
		if len(line) == 0 {
			continue
		}
		l := string(line)
		switch strings.Fields(l)[0] {
		case "dump":
			io.WriteString(c, fmt.Sprintf("Users: %v\r\n", d.s.Users))
			io.WriteString(c, fmt.Sprintf("Channels: %v\r\n", d.s.Channels))
		default:
			d.s.sendq <- l
		}
	}
}
