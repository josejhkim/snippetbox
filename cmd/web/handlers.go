package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home handler function for writing a byte slice.
func home(w http.ResponseWriter, r *http.Request) {
	// Uses the Header().Add() function to add a "Server: Go" header
	// to the response header map.
	// Any header map modifications must be done before calling
	// w.WriteHeader()
	w.Header().Add("Server", "Go")

	// http's DetectContentType() automatically detects the content type
	// and sets the content type accordingly in the header
	// but one caveat is that it can't distinguish application/json from text/plain

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// All wildcard values are read in Golang as string
	// so we must do any sanitization / data format conversion ourselves.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	/*
		msg := fmt.Sprintf("Display a specific snippet with ID %d", id)
		w.Write([]byte(msg))
	*/
	// Because http.ResponseWriter satisfies the io.Writer interface
	// we can do this instead:
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new request"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Writes a different status code in the response header.
	// WriteHeader() can only be called once per response.
	// If WriteHeader() isn't called explicitly, then 200 will automatically be sent
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}
