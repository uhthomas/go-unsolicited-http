package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Main(ctx context.Context) error {
	for i := 0; ; i++ {
		fmt.Println("attempt", i)

		if _, err := http.Head("http://localhost:8080"); err != nil {
			return fmt.Errorf("head: %w", err)
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
