package utils

import (
	"os"
	"os/signal"
)

// listener to get signal from goroutine
func BlockUntilSignal(signals ...os.Signal) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, signals...)
	<-done
}
