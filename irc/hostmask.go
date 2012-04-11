package irc

import (
	"fmt"
	"strings"
)

type HostMask struct {
	Nick, Ident, Host string
}

func (h HostMask) String() string {
	return fmt.Sprintf("%s!%s@%s", h.Nick, h.Ident, h.Host)
}

func Parse(str string) *HostMask {
	hm := new(HostMask)
	hm.Host = str
	nidx, uidx := strings.Index(str, "!"), strings.Index(str, "@")
	if nidx != -1 && uidx != -1 {
		hm.Nick = str[:nidx]
		hm.Ident = str[nidx+1 : uidx]
		hm.Host = str[uidx+1:]
	}
	return hm
}
