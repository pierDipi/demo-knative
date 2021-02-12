package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func functionContext() context.Context {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigs
		cancel()
	}()

	return ctx
}
