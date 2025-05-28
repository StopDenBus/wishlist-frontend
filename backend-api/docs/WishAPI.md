# \WishAPI

All URIs are relative to *https://wishlist-backend.apps.gusek.info*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWishWishPost**](WishAPI.md#CreateWishWishPost) | **Post** /wish | Create Wish
[**DeleteWishByIdWishByIdIdDelete**](WishAPI.md#DeleteWishByIdWishByIdIdDelete) | **Delete** /wish/by_id/{id} | Delete Wish By Id
[**DeleteWishByNameWishByNameWishDelete**](WishAPI.md#DeleteWishByNameWishByNameWishDelete) | **Delete** /wish/by_name/{wish} | Delete Wish By Name
[**GetWishByIdWishByIdIdGet**](WishAPI.md#GetWishByIdWishByIdIdGet) | **Get** /wish/by_id/{id} | Get Wish By Id
[**GetWishByNameWishByNameWishGet**](WishAPI.md#GetWishByNameWishByNameWishGet) | **Get** /wish/by_name/{wish} | Get Wish By Name
[**GetWishesWishesGet**](WishAPI.md#GetWishesWishesGet) | **Get** /wishes | Get Wishes
[**UpdateWishWishIdPut**](WishAPI.md#UpdateWishWishIdPut) | **Put** /wish/{id} | Update Wish



## CreateWishWishPost

> Wish CreateWishWishPost(ctx).WishIn(wishIn).Execute()

Create Wish



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	wishIn := *openapiclient.NewWishIn("Product_example", float32(123), "Url_example", openapiclient.PriorityEnum("Hoch")) // WishIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.CreateWishWishPost(context.Background()).WishIn(wishIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.CreateWishWishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWishWishPost`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.CreateWishWishPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWishWishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wishIn** | [**WishIn**](WishIn.md) |  | 

### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWishByIdWishByIdIdDelete

> Wish DeleteWishByIdWishByIdIdDelete(ctx, id).Execute()

Delete Wish By Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.DeleteWishByIdWishByIdIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.DeleteWishByIdWishByIdIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWishByIdWishByIdIdDelete`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.DeleteWishByIdWishByIdIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWishByIdWishByIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWishByNameWishByNameWishDelete

> Wish DeleteWishByNameWishByNameWishDelete(ctx, wish).Execute()

Delete Wish By Name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	wish := "wish_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.DeleteWishByNameWishByNameWishDelete(context.Background(), wish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.DeleteWishByNameWishByNameWishDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWishByNameWishByNameWishDelete`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.DeleteWishByNameWishByNameWishDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wish** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWishByNameWishByNameWishDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWishByIdWishByIdIdGet

> Wish GetWishByIdWishByIdIdGet(ctx, id).Execute()

Get Wish By Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.GetWishByIdWishByIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.GetWishByIdWishByIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWishByIdWishByIdIdGet`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.GetWishByIdWishByIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWishByIdWishByIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWishByNameWishByNameWishGet

> Wish GetWishByNameWishByNameWishGet(ctx, wish).Execute()

Get Wish By Name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	wish := "wish_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.GetWishByNameWishByNameWishGet(context.Background(), wish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.GetWishByNameWishByNameWishGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWishByNameWishByNameWishGet`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.GetWishByNameWishByNameWishGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wish** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWishByNameWishByNameWishGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWishesWishesGet

> []Wish GetWishesWishesGet(ctx).SortBy(sortBy).OrderBy(orderBy).Execute()

Get Wishes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sortBy := openapiclient.SortBy("product") // SortBy |  (optional)
	orderBy := openapiclient.OrderBy("ascending") // OrderBy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.GetWishesWishesGet(context.Background()).SortBy(sortBy).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.GetWishesWishesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWishesWishesGet`: []Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.GetWishesWishesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWishesWishesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | [**SortBy**](SortBy.md) |  | 
 **orderBy** | [**OrderBy**](OrderBy.md) |  | 

### Return type

[**[]Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWishWishIdPut

> Wish UpdateWishWishIdPut(ctx, id).WishIn(wishIn).Execute()

Update Wish



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | 
	wishIn := *openapiclient.NewWishIn("Product_example", float32(123), "Url_example", openapiclient.PriorityEnum("Hoch")) // WishIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishAPI.UpdateWishWishIdPut(context.Background(), id).WishIn(wishIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishAPI.UpdateWishWishIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWishWishIdPut`: Wish
	fmt.Fprintf(os.Stdout, "Response from `WishAPI.UpdateWishWishIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWishWishIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wishIn** | [**WishIn**](WishIn.md) |  | 

### Return type

[**Wish**](Wish.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

