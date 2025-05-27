package wish

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/StopDenBus/wishlist-frontend/utils"
)

type Handler struct{}

type Wish struct {
	ID       int     `json:"id"`
	Product  string  `json:"product"`
	Price    float32 `json:"price"`
	Priority string  `json:"priority"`
	Url      string  `json:"url"`
}

var wishes []Wish

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", getWishes)
	r.HandleFunc("POST /", createWish)

	r.HandleFunc("GET /{id}/", getWish)
	r.HandleFunc("DELETE /{id}/", deleteWish)

	return r
}

func getWishes(w http.ResponseWriter, r *http.Request) {
	requestURL := "https://wishlist-backend.apps.gusek.info/wishes"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	err = json.NewDecoder(r.Body).Decode(&wishes)
	if err != nil {
		fmt.Printf("error parsing request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Println(wishes)
	// response := map[string]any{
	// 	"message": "Done",
	// 	"wishes":  wishes,
	// }
	// utils.WriteJSONResponse(w, http.StatusOK, response)
}

func getWish(w http.ResponseWriter, r *http.Request) {
	wishID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response := map[string]any{
			"message": "Invalid wish id. Wish id must be of type int.",
		}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	for _, wish := range wishes {
		if wish.ID == wishID {
			response := map[string]any{
				"message": "Done",
				"wish":    wish,
			}
			utils.WriteJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	response := map[string]any{"message": fmt.Sprintf("Wish with id: %d not found", wishID)}
	utils.WriteJSONResponse(w, http.StatusNotFound, response)
}

func createWish(w http.ResponseWriter, r *http.Request) {
	var wish Wish
	if err := utils.ParseJSON(r, &wish); err != nil {
		response := map[string]any{"message": "Invalid request"}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	wishes = append(wishes, wish)
	response := map[string]any{"message": "Done"}
	utils.WriteJSONResponse(w, http.StatusCreated, response)
}

func deleteWish(w http.ResponseWriter, r *http.Request) {
	wishID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response := map[string]any{
			"message": "Invalid wish id. Wish id must be of type int.",
		}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	wishIndex := -1
	for index, wish := range wishes {
		if wishID == wish.ID {
			wishIndex = index
			break
		}
	}

	if wishIndex == -1 {
		response := map[string]any{"message": fmt.Sprintf("Wish with id: %d not found", wishID)}
		utils.WriteJSONResponse(w, http.StatusNotFound, response)
		return
	}

	wishes = append(wishes[:wishIndex], wishes[wishIndex+1:]...)
	response := map[string]any{"message": fmt.Sprintf("Deleted wish with id: %d", wishID)}
	utils.WriteJSONResponse(w, http.StatusNotFound, response)
}
