package wish

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	backend "github.com/StopDenBus/wishlist-frontend/backend"
	"github.com/StopDenBus/wishlist-frontend/utils"
)

var ctx = context.WithValue(context.Background(), backend.ContextServerIndex, 1)

type Handler struct{}

type Wish struct {
	ID       int     `json:"id"`
	Product  string  `json:"product"`
	Price    float32 `json:"price"`
	Priority string  `json:"priority"`
	Url      string  `json:"url"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", getWishes)
	r.HandleFunc("POST /", createWish)

	r.HandleFunc("GET /by_id/{id}/", getWish)
	r.HandleFunc("DELETE /by_id/{id}/", deleteWish)
	r.HandleFunc("PUT /{id}/", updateWish)

	return r
}

func getWishes(w http.ResponseWriter, r *http.Request) {
	sortBy := backend.SortBy("product")     // SortBy |  (optional)
	orderBy := backend.OrderBy("ascending") // OrderBy |  (optional)

	u, _ := url.Parse(r.URL.RequestURI())
	q := u.Query()
	order := q.Get("order_by")
	if order != "ascending" {
		orderBy = backend.OrderBy("descending")
	}

	sort := q.Get("sort_by")
	if sort != "" {
		sortBy = backend.SortBy(sort)
	}

	configuration := backend.NewConfiguration()
	apiClient := backend.NewAPIClient(configuration)
	wishes, response, err := apiClient.WishAPI.GetWishesWishesGet(ctx).SortBy(sortBy).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.GetWishesWishesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", response)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Got error status code: %v\n", response)
	}
	// response from `GetWishesWishesGet`: []Wish
	// fmt.Fprintf(os.Stdout, "Response from `WishAPI.GetWishesWishesGet`: %v\n", wishes)

	utils.WriteJSONResponse(w, http.StatusOK, wishes)
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

	configuration := backend.NewConfiguration()
	apiClient := backend.NewAPIClient(configuration)
	wish, apiResponse, err := apiClient.WishAPI.GetWishByIdWishByIdIdGet(ctx, int32(wishID)).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.GetWishByIdWishByIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	utils.WriteJSONResponse(w, apiResponse.StatusCode, wish)

}

func createWish(w http.ResponseWriter, r *http.Request) {
	var wish Wish
	if err := utils.ParseJSON(r, &wish); err != nil {
		response := map[string]any{"message": "Invalid request"}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	wishIn := *backend.NewWishIn(wish.Product, float32(wish.Price), wish.Url, backend.PriorityEnum(wish.Priority))

	configuration := backend.NewConfiguration()
	apiClient := backend.NewAPIClient(configuration)
	_, apiResponse, err := apiClient.WishAPI.CreateWishWishPost(ctx).WishIn(wishIn).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.CreateWishWishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	utils.WriteJSONResponse(w, apiResponse.StatusCode, apiResponse.Body)
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

	configuration := backend.NewConfiguration()
	apiClient := backend.NewAPIClient(configuration)
	_, apiResponse, err := apiClient.WishAPI.DeleteWishByIdWishByIdIdDelete(ctx, int32(wishID)).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.DeleteWishByIdWishByIdIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	utils.WriteJSONResponse(w, apiResponse.StatusCode, apiResponse.Body)
}

func updateWish(w http.ResponseWriter, r *http.Request) {
	var wish Wish

	if err := utils.ParseJSON(r, &wish); err != nil {
		response := map[string]any{"message": "Invalid request"}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	wishID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		response := map[string]any{
			"message": "Invalid wish id. Wish id must be of type int.",
		}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	wishIn := *backend.NewWishInWithDefaults()
	wishIn.SetPrice(wish.Price)
	wishIn.SetPriority(backend.PriorityEnum(wish.Priority))
	wishIn.SetProduct(wish.Product)
	wishIn.SetUrl(wish.Url)

	configuration := backend.NewConfiguration()
	apiClient := backend.NewAPIClient(configuration)

	_, apiResponse, err := apiClient.WishAPI.UpdateWishWishIdPut(ctx, int32(wishID)).WishIn(wishIn).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.UpdateWishWishIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	utils.WriteJSONResponse(w, apiResponse.StatusCode, apiResponse.Body)
}
