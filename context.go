package tangent

import (
	"net"
)

// Context represents a tracking instance
type Context struct {
	servers []*Server    // Servers we know of
	events  chan Event   // Event dispatch channel
	dead    chan *Server // Death notification channel
	conns   int          // Active connections
	Events  <-chan Event // Event API
}

// New creates a new Context ready to be used
func New() *Context {
	ev := make(chan Event, 10)
	c := &Context{events: ev, Events: ev, dead: make(chan *Server, 5)}
	go c.monitor()
	return c
}

// Connect creates a new IRC connection
func (c *Context) Connect(addr string, ssl bool, user *User) (srv *Server, err error) {
	//TODO: SSL
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Connected to ", addr)
	srv = &Server{conn: nc, connected: true, context: c, us: user}
	c.servers = append(c.servers, srv)
	srv.start()
	return
}

func (c *Context) Stop(msg string) {
	for _, s := range c.servers {
		if s.connected {
			s.Quit(msg)
		}
	}
}

func (c *Context) monitor() {
	for death := range c.dead {
		for i, s := range c.servers {
			if s == death {
				c.servers = append(c.servers[:i], c.servers[i+1:]...)
			}
		}
		c.conns--
		if c.conns <= 0 {
			close(c.dead)
			close(c.events)
		} else {
			logger.Debug("Conns is now ", c.conns)
		}
	}
}
