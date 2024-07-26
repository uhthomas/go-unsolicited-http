package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Main(ctx context.Context) error {
	c := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
		},
	}

	for i := 0; ; i++ {
		fmt.Println("attempt", i)

		req, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost:8080", nil)
		if err != nil {
			return fmt.Errorf("new request: %w", err)
		}

		if _, err := c.Do(req); err != nil {
			return fmt.Errorf("do: %w", err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
