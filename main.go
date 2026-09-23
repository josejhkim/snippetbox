package main

import (
	"fmt"
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

func main() {
	// ServeMux functions as a router that maps the request URLs
	// to the corresponding handlers.
	mux := http.NewServeMux()

	// Registers the home function as the handler for the "/" URL pattern
	// In Golang's servemux, route paths trailing slashes act as a wildcard,
	// and handle every request route for its recursive subtrees.
	// [eg. mux.HandleFunc("/", home) ]
	// whereas route paths with no trailing slashes only handle the specified route
	// to restrict the wildcard matching for trailing slashes, we can attach a {$}.

	// Prefixing the route with the method makes it only handle the matching method
	// for the given route.
	// GET also matches HEAD, which is used to retrieve the headers of a resource
	// without actually downloading the resource
	mux.HandleFunc("GET /{$}", home)

	// If this route was enabled instead of /snippet/create, then any requests to /snippet/create
	// would be redirected here instead.
	// mux.HandleFunc("/snippet/create/", snippetCreate2)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// Okay to register 2 same routes under different methods
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Registers an {id} wildcard.
	// All wildcards are read as strings.

	// If we have something like "snippet/view/{id}/" instead,
	// the rules about trailing slash also applies here in that
	// /snippet/view/2/a/b/c will still be matched with this route handler.
	// But this will only return 2 as the path param.
	// If we want the entire subtree as the path param,
	// we can do "snippet/view/{id...}"
	// to get the entire trailing values as a path param.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)

	// If we have an overlapping route like
	// mux.HandleFunc("/snippet/view/edit", snippetEdit)
	// then the more specific route pattern wins.

	// For the example above, the "/snippet/view/edit" route wins
	// because it's more specific in the sense
	// that the word 'edit' is included in all possible {id} path params.

	// As another example, if we have 2 routes
	// /post/new/{id}
	// /post/{author}/latest
	// then this causes an error.

	// Prints a log message to say the server is starting.
	log.Print("starting server on :4000")

	// http.ListenAndServe blocks until it stops executing
	// in which case it will return an error, which will in turn
	// be logged via log.Fatal().
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	/*
		It's possible to do something like
		http.HandleFunc("/home", homeHandler)
		and then
		http.ListenAndServe(":4000", nil)

		This registers the routes under http.DefaultServeMux, which is a global variable in the Golang STL.
		This isn't recommended because
		1) It's hard to see where your routes and handlers are declared/attached to
		2) Compromised third-party libraries can expose malicious routes on DefaultServeMux
	*/
}
