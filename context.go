package tangent

import (
	"net"
)

// Context represents a tracking instance
type Context struct {
	servers []*Server // Servers we know of
}

// New creates a new Context ready to be used
func New() *Context {
	return &Context{}
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
