package calc

import (
	"log"
	"net/http"
	"time"
)

// Start launches calc service.
func Start(address string) error {
	mux := http.NewServeMux()

	for route, handler := range Routes {
		mux.Handle(route, handler)
	}

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("calc service is listening at %s", address)
	return s.ListenAndServe()
}
