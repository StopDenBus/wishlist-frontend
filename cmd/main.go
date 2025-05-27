package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/StopDenBus/wishlist-frontend/api"
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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pic/favicon.ico")
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
