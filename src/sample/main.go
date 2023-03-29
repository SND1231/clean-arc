package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	r := getRouter()
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("go testだよな")
	go func() {
		defer wg.Done()
		http.ListenAndServe(":3010", r)
	}()
	waitSignal()
	wg.Wait()
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/test", test)
	return r
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start")
	time.Sleep(time.Second * 10)
	fmt.Println("End")
	w.WriteHeader(http.StatusOK)
}

func waitSignal() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)
	<-sigChan
}
