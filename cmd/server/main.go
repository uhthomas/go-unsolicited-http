package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Server struct {
	requestCount uint64
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := atomic.AddUint64(&s.requestCount, 1)

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	res, err := http.Get("https://6f.io")
	if err != nil {
		panic(err)
	}

	n, err := io.Copy(w, io.LimitReader(res.Body, 54))

	fmt.Printf("%d: %d bytes written, err=%v\n", requestID, n, err)
}

func Main(ctx context.Context) error {
	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()
	return (&http.Server{
		Addr:         *addr,
		Handler:      &Server{},
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}).ListenAndServe()
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
