package tangent

import (
    "github.com/xales/zephyr/log"
    "os"
)

var logger *log.Logger

// Server represents a server on the network and all our tracking data about it
// It also handles the network connection to the server
type Server struct {
}


func init() {
    logger = log.New(os.Stdout, log.Debug0, "[IRC] ")
}
