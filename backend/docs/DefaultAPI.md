# \DefaultAPI

All URIs are relative to *https://wishlist-backend.apps.gusek.info*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadinessReadinessGet**](DefaultAPI.md#ReadinessReadinessGet) | **Get** /readiness | Readiness



## ReadinessReadinessGet

> interface{} ReadinessReadinessGet(ctx).Execute()

Readiness

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReadinessReadinessGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReadinessReadinessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadinessReadinessGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReadinessReadinessGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadinessReadinessGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

