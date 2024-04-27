package main

import (
	"context"
	"github.com/alexisleon/stori/cmd"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	execCtx, execCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer execCancel()

	if err := cmd.RootCmd().ExecuteContext(execCtx); err != nil {
		log.Fatal(err)
	}
}
