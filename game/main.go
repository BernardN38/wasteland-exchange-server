package main

import "net/http"

func main() {
	// This is a placeholder for the main function.
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! \n"))
	}))
}
