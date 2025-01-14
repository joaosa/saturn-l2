package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var port int
	portStr := os.Getenv("PORT")
	if portStr == "" {
		port = 5500
	} else {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			panic(fmt.Errorf("Invalid PORT value '%s': %s", portStr, err.Error()))
		}
	}

	m := mux.NewRouter()
	m.Handle("/webui", http.HandlerFunc(webuiIndex))
	srv := &http.Server{
		Handler: m,
	}

	nl, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := srv.Serve(nl); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	port = nl.Addr().(*net.TCPAddr).Port
	fmt.Println("Server listening on", nl.Addr())
	fmt.Printf("WebUI: http://localhost:%d/webui\n", port)
	for {

	}
}

func webuiIndex(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<html><head><title>Saturn L2 Node</title></head><body>Status: running</body></html>")
}
