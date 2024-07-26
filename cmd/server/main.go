package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type Server struct {
	requestCount uint64
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := atomic.AddUint64(&s.requestCount, 1)

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	res, err := http.Get("https://go.dev")
	if err != nil {
		panic(err)
	}

	n, err := io.Copy(w, res.Body)

	fmt.Printf("%d: %d bytes written, err=%v\n", requestID, n, err)
}

func Main(ctx context.Context) error {
	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()

	log.Println("listening on", *addr)

	return http.ListenAndServe(*addr, &Server{})
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
