// Command sdnotifytest is an integration test command which sends systemd
// readiness notifications to a socket and via stdout.
package main

import (
	"log"

	"github.com/mdlayher/sdnotify"
)

func main() {
	log.SetFlags(0)

	n, err := sdnotify.New()
	if err != nil {
		log.Fatalf("failed to open notifier: %v", err)
	}
	defer n.Close()

	for i := 0; i < 3; i++ {
		out(n, sdnotify.Statusf("waiting %d", i))
	}

	out(n, sdnotify.Ready, sdnotify.Statusf("done"), sdnotify.Stopping)
}

func out(n *sdnotify.Notifier, ss ...string) {
	if err := n.Notify(ss...); err != nil {
		log.Fatalf("failed to notify: %v", err)
	}
}
