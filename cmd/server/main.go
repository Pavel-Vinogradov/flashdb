package main

import (
	"context"
	"flashdb/cmd/cli"
	"flashdb/internal/server"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv, err := server.NewServer("8080")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := srv.Start(); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	app := cli.NewApp()
	app.RegisterCommand("ping", func(args []string) error {
		_, err := conn.Write([]byte("PING\n"))
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}

		fmt.Println(string(buf[:n]))
		return nil
	})
	app.Run()

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("Shutdown signal received")
		cancel()
		srv.Stop()
	}()

	<-ctx.Done()
	fmt.Println("Server stopped")
}
