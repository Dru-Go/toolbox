package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dru-go/noah-toolbox/usecase"
	"github.com/gorilla/mux"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func SetupRouter(MaterialUsecase usecase.IMaterialUsecase, TransactionUsecase usecase.ITransactionUsecase) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Still Alive")
	})
	m := router.PathPrefix("/material").Subrouter()
	m.HandleFunc("/import", ImportHandler(MaterialUsecase)).Methods("POST")
	t := router.PathPrefix("/transaction").Subrouter()
	t.HandleFunc("/import", ImportTransactionHandler(TransactionUsecase)).Methods("POST")
	t.HandleFunc("/compute", ComputeTransactionHandler(TransactionUsecase)).Methods("POST")
	srv := &http.Server{
		Addr: "0.0.0.0:3200",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
