// Package tangent is an IRC library with advanced tracking capabilities
package tangent

import (
	"bufio"
	"fmt"
	"github.com/xales/Zephyr/log"
	"github.com/xales/tangent/irc"
	"net"
	"os"
	"strings"
)

var logger *log.Logger

// Server represents a server on the network and all our tracking data about it
// It also handles the network connection to the server
type Server struct {
	conn      net.Conn            // Network connection
	connected bool                // State of connection
	buf       *bufio.ReadWriter   // Buffered IO
	sendq     chan string         // Send queue
	endq      chan bool           // Semaphore for writer
	Name      string              // Server's name
	context   *Context            // Context the server exists in
	us        *User               // User instance for our connection
	capabs    map[string]bool     // Capabilities enabled
	Channels  map[string]*Channel // Channels
	Users     map[string]*User    // Users
}

func (s *Server) FindUser(nick string) *User {
	return s.Users[strings.ToLower(nick)]
}

func (s *Server) FindChan(chn string) *Channel {
	return s.Channels[strings.ToLower(chn)]
}

func (s *Server) fuser(hm *irc.HostMask) *User {
	if u := s.FindUser(hm.Nick); u != nil {
		return u
	}
	u := &User{HostMask: *hm, Channels: make(map[*Channel]bool)}
	s.Users[strings.ToLower(hm.Nick)] = u
	return u
}

func (s *Server) fchan(name string) *Channel {
	if ch := s.FindChan(name); ch != nil {
		return ch
	}
	ch := &Channel{Name: name, Users: make(map[*User]bool)}
	s.Channels[strings.ToLower(name)] = ch
	return ch
}

// Initializes structures and starts goroutines to handle networking, and sends login info
func (s *Server) start() {
	s.us.Channels = make(map[*Channel]bool)
	s.Channels = make(map[string]*Channel)
	s.Users = make(map[string]*User)
	s.capabs = make(map[string]bool)
	s.buf = bufio.NewReadWriter(bufio.NewReader(s.conn), bufio.NewWriter(s.conn))
	s.sendq = make(chan string)
	s.endq = make(chan bool, 1)
	s.connected = true
	go s.reader()
	go s.writer()
	s.RawCmd("CAP", "LS")
}

func (s *Server) register() {
	s.Cmd("USER", s.us.Ident, "*", "*", s.us.Real)
	s.Nick(s.us.Nick)
}

func (s *Server) reader() {
	for {
		line, prefix, err := s.buf.ReadLine()
		if err != nil {
			logger.Error(err)
			//TODO allow reconnect
			close(s.sendq)
			<-s.endq
			close(s.endq)
			s.connected = false
			s.context.dead <- s
			return
		}
		if prefix {
			logger.Warning("Network overflow; ignoring line")
			continue
		}
		logger.Debugf("-> %s", line)
		if len(line) == 0 {
			continue
		}
		l := irc.ParseLine(string(line))
		if l != nil {
			s.handle(l)
		}
	}
}

func (s *Server) handle(line *irc.Line) {
	if cmd, ok := scommands[line.Cmd]; ok {
		if len(line.Args) > cmd.args {
			cmd.handle(line, s)
		}
	}
}

func (s *Server) writer() {
	for line := range s.sendq {
		logger.Debugf("<- %s", line)
		_, err := s.buf.WriteString(line + "\r\n")
		if err != nil {
			logger.Error(err)
			break
		}
		err = s.buf.Flush() // Do we want this?
		if err != nil {
			logger.Error(err)
			break
		}
	}
	s.endq <- true
}

// Cmd sends a command to the server, prefixing the final argument with :
func (s *Server) Cmd(cmd string, args ...string) {
	if len(args) > 0 {
		args[len(args)-1] = ":" + args[len(args)-1]
	}
	s.sendq <- fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
}

// Cmd sends a command to the server, sending arguments as-is
func (s *Server) RawCmd(cmd string, args ...string) {
	s.sendq <- fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
}

func (s *Server) Nick(nick string) {
	s.RawCmd("NICK", nick)
}

func (s *Server) e(ev Event) {
	ev.setSource(s)
	s.context.events <- ev
}

func (s *Server) Quit(msg string) {
	s.Cmd("QUIT", msg)
}

func init() {
	logger = log.New(os.Stdout, log.Debug0, "[IRC] ")
}
