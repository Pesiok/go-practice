package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func runContext() {
	http.HandleFunc("/", fooOne)
	http.HandleFunc("/bar", barOne)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fooOne(w http.ResponseWriter, req *http.Request) {
	// get info about request context
	ctx := req.Context()

	// only values directly associated with request
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// ~~ ridiculous long running task
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	// get value from the channel that will return value first
	select {
	// value from cancel()
	case <-ctx.Done():
		// stops gouroutine
		return 0, ctx.Err()
	// value from gouroutine (long task)
	case i := <-ch:
		return i, nil
	}
}

func barOne(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
