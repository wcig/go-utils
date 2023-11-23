package xsignal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Notify(sig ...os.Signal) {
	if len(sig) == 0 {
		panic("notify signal empty")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, sig...)
	quitSignal := <-quit
	log.Println(">> quit:", quitSignal)
}

func NotifyDefault() {
	// syscall.SIGINT: Ctrl+C, syscall.SIGTERM: Kill
	Notify(syscall.SIGINT, syscall.SIGTERM)
}
