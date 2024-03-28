package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

// Listen- blocking call which listens for a shutdown signal and calls the function f
func Listen(ctx context.Context, f func()) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			f()
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	return ctx, cancel
}
