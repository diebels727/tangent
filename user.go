package tangent

import (
	"fmt"
	"github.com/xales/tangent/irc"
	"strings"
)

// User represents information about a user on IRC
type User struct {
	irc.HostMask
	Real     string            // Realname(GECOS)
	Account  string            // Services account name
	Channels map[*Channel]bool // Channels user is in
}

func (u *User) String() string {
	chn := make([]string, 0, len(u.Channels))
	for ch := range u.Channels {
		chn = append(chn, ch.Name)
	}
	return fmt.Sprintf("[User %s (acc: %s) | %s]", u.HostMask, u.Account, strings.Join(chn, " "))
}

type Channel struct {
	Name  string         // Name of channel
	Users map[*User]bool // Users in channel
}

func (c *Channel) String() string {
    usr := make([]string, 0, len(c.Users))
    for u := range c.Users {
        usr = append(usr, u.Nick)
    }
	return fmt.Sprintf("[Chan %s | %v]", c.Name, strings.Join(usr, " "))
}
