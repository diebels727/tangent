package tangent

// An Event is a processed event for handling by user routines
type Event interface {
	Source() *Server
	setSource(*Server)
}

type defev struct {
	src *Server
}

func (d *defev) Source() *Server {
	return d.src
}

func (d *defev) setSource(s *Server) {
	d.src = s
}

type Connect struct {
	defev
}
