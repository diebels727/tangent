package main

import (
	"fmt"
	"github.com/xales/tangent"
	"github.com/xales/tangent/irc"
	"os"
	"os/signal"
)

func main() {
	c := tangent.New()
	u := &tangent.User{HostMask: irc.HostMask{Nick: "xales", Ident: "ttest"}, Real: "ttest"}
	s, err := c.Connect("irc.esper.net:6667", false, u)
	if err != nil {
		panic(err)
	}
	_, _ = s.Debug(":4444")
	go func() {
		chm := make(chan os.Signal)
		signal.Notify(chm, os.Interrupt)
		<-chm
		c.Stop("Interrupt")
	}()
	for ev := range c.Events {
		fmt.Printf("%#v\n", ev)
	}
}
