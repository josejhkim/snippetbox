package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home handler function for writing a byte slice.
func home(w http.ResponseWriter, r *http.Request) {
	// Uses the Header().Add() function to add a "Server: Go" header
	// to the response header map.
	// Any header map modifications must be done before calling
	// w.WriteHeader()

	// http's DetectContentType() automatically detects the content type
	// and sets the content type accordingly in the header
	// but one caveat is that it can't distinguish application/json from text/plain
	w.Header().Add("Server", "Go")
	// ---

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file
	// into a template set.
	// The file path must either be relative to the CWD, or an absolute path
	// ts, err := template.ParseFiles("./ui/html/base.tmpl")
	// ---
	// Files can also be passed as variadic arguments
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Use the Execute() method on the template set to write the template content
	// into the response body.
	// The last parameter represents dynamic data that we want to pass in.
	// err = ts.Execute(w, nil)
	// ---
	// Uses the ExecuteTemplate() method to write the contents of the "base"
	// template as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
