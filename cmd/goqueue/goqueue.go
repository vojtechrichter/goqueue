package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/vojtechrichter/goqueue/internal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	_, err := internal.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello from goqueue!")
}
