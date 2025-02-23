package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/olivere/vite"
)

//go:embed all:dist
var dist embed.FS

func main() {
	var (
		isDev = flag.Bool("dev", false, "run in development mode")
	)
	flag.Parse()

	if *isDev {
		runDevServer()
	} else {
		runProdServer()
	}
}

func runDevServer() {
	// Handle the Vite server.
	viteHandler, err := vite.NewHandler(vite.Config{
		FS:      os.DirFS("."),
		IsDev:   true,
		ViteURL: "http://localhost:5173",
	})
	if err != nil {
		panic(err)
	}

	// Create a new handler.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			// Server the index.html file.
			ctx := r.Context()
			ctx = vite.MetadataToContext(ctx, vite.Metadata{
				Title: "Hello, Vite!",
			})
			ctx = vite.ScriptsToContext(ctx, `<script>console.log('Hello, nice to meet you in the console!')</script>`)
			viteHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		viteHandler.ServeHTTP(w, r)
	})

	// Start a listener.
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		var err1 error
		if l, err1 = net.Listen("tcp6", "[::1]:0"); err1 != nil {
			panic(fmt.Errorf("starting HTTP server: %v", err))
		}
	}

	// Create a new server.
	server := http.Server{
		Handler: handler,
	}

	log.Printf("Listening on on http://%s", l.Addr())

	// Start the server.
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}

type Response struct {
	Data string `json:data`
}

func runProdServer() {
	fs, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	// Create a new handler.
	viteHandler, err := vite.NewHandler(vite.Config{
		FS:    fs,
		IsDev: false,
	})
	if err != nil {
		panic(err)
	}

	// Create a new handler.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			// Server the index.html file.
			ctx := r.Context()
			ctx = vite.MetadataToContext(ctx, vite.Metadata{
				Title: "Hello, Vite!",
			})
			ctx = vite.ScriptsToContext(ctx, `<script>console.log('Hello, nice to meet you in the console!')</script>`)
			viteHandler.ServeHTTP(w, r.WithContext(ctx))
			return
		} else if r.URL.Path == "/api/test" && r.Method == "GET" {
			res := Response{
				Data: "This is response",
			}

			parsed, err := json.Marshal(res)
			if err != nil {
				return
			}

			w.Write(parsed)
			// w.Header().Set("mimetype", "application/json")
			// w.
			return
		}

		viteHandler.ServeHTTP(w, r)
	})
	// Start a listener.
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		var err1 error
		if l, err1 = net.Listen("tcp6", "[::1]:0"); err1 != nil {
			panic(fmt.Errorf("starting HTTP server: %v", err))
		}
	}

	// Create a new server.
	server := http.Server{
		Handler: handler,
	}

	log.Printf("Listening on on http://%s", l.Addr())

	// Start the server.
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}
