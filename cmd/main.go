package main

import (
	"context"
	"fmt"
	"github.com/ninadingole/bazel-for-go/pkg"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Hello, world.")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill)

	mux := http.NewServeMux()
	mux.Handle("/", pkg.HelloHandler())

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		<-shutdown
		fmt.Println("\nShutting down server...")
		timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()

		err := server.Shutdown(timeout)
		if err != nil {
			fmt.Println("Error shutting down server:", err)
		}
	}()

	fmt.Println("Server is listening on port 8080")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Error starting server:", err)
	}

}
