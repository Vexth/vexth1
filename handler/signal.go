package handler

import (
	"os/signal"
	"syscall"
)

func (ep *Entrypoint) configureSignals() {
	// signal.Notify(ep.signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	signal.Notify(ep.signals, syscall.SIGINT, syscall.SIGTERM)
}

func (ep *Entrypoint) listenSignals() {
	for {
		sig := <-ep.signals

		switch sig {
		// case syscall.SIGUSR1:
		default:
			ep.Stop()
			return
		}
	}
}
