package main

import (
	"log"

	"github.com/StopDenBus/wishlist-frontend/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	// // Handle favicon
	// http.HandleFunc("/favicon.ico", faviconHandler)
	// http.HandleFunc("/", serveTemplate)
	// // Handle client's requests for CSS
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	// // Handle client's requests for JS
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	// log.Println("Server starting on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatalf("could not start server: %s\n", err)
	// }
}
