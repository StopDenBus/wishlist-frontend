package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/StopDenBus/wishlist-frontend/services/wish"
	"github.com/StopDenBus/wishlist-frontend/utils"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	wishHandler := wish.NewHandler()
	wishRouter := wishHandler.RegisterRoutes()

	router := http.NewServeMux()
	router.Handle("/wish/", http.StripPrefix("/wish", wishRouter))
	// Handle favicon
	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/", serveTemplate)
	// Handle client's requests for CSS
	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	// Handle client's requests for JS
	router.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	// Handle readiness
	router.HandleFunc("/readiness", readinessHandler)

	log.Println("Starting server on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pic/favicon.ico")
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{"status": "ok"}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
