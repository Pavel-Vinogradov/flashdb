package main

import (
	"context"
	"flashdb/cmd/cli"
	"flashdb/internal/server"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
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

		fmt.Print(string(buf[:n]))
		return nil
	})
	app.RegisterCommand("set", func(args []string) error {
		if len(args) < 2 {
			fmt.Println("ERR wrong number of arguments for SET")
			return nil
		}
		cmd := "SET " + args[0] + " " + strings.Join(args[1:], " ") + "\n"
		_, err := conn.Write([]byte(cmd))
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}

		fmt.Print(string(buf[:n]))
		return nil
	})
	app.RegisterCommand("get", func(args []string) error {
		if len(args) < 1 {
			fmt.Println("ERR wrong number of arguments for GET")
			return nil
		}
		cmd := "GET " + args[0] + "\n"
		_, err := conn.Write([]byte(cmd))
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}

		fmt.Print(string(buf[:n]))
		return nil
	})
	app.RegisterCommand("del", func(args []string) error {
		if len(args) < 1 {
			fmt.Println("ERR wrong number of arguments for DEL")
			return nil
		}
		cmd := "DEL " + args[0] + "\n"
		_, err := conn.Write([]byte(cmd))
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}

		fmt.Print(string(buf[:n]))
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
