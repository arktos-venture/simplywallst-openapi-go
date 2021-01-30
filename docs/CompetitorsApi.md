# \CompetitorsApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCompetitors**](CompetitorsApi.md#ListCompetitors) | **Get** /api/competitors/{id} | 



## ListCompetitors

> Competitors ListCompetitors(ctx, id).Include(include).Version(version).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "["C163BA6A-EA32-4730-99B2-376697240A73"]" // string | string id (name or id) of the competitors
    include := "["info","score","analysis.raw_data","score.snowflake","analysis.extended.raw_data","analysis.extended.raw_data.insider_transactions","analysis.raw_data.insider_transactions"]" // string | string include (name or id) of the competitors (optional)
    version := "[2]" // string | string version (name or id) of the competitors (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompetitorsApi.ListCompetitors(context.Background(), id).Include(include).Version(version).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompetitorsApi.ListCompetitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompetitors`: Competitors
    fmt.Fprintf(os.Stdout, "Response from `CompetitorsApi.ListCompetitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | string id (name or id) of the competitors | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompetitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | string include (name or id) of the competitors | 
 **version** | **string** | string version (name or id) of the competitors | 

### Return type

[**Competitors**](competitors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

