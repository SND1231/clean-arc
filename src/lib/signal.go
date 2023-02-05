package lib

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
	<-sigChan
}
