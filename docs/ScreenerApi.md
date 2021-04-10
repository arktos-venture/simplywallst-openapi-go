# \ScreenerApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGridViewSearch**](ScreenerApi.md#ListGridViewSearch) | **Post** /api/grid/view/search | 
[**ListGrids**](ScreenerApi.md#ListGrids) | **Post** /api/grid/filter | 



## ListGridViewSearch

> GridViewSearch ListGridViewSearch(ctx).GridViewSearchAdd(gridViewSearchAdd).Execute()





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
    gridViewSearchAdd := *openapiclient.NewGridViewSearchAdd() // GridViewSearchAdd |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScreenerApi.ListGridViewSearch(context.Background()).GridViewSearchAdd(gridViewSearchAdd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScreenerApi.ListGridViewSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGridViewSearch`: GridViewSearch
    fmt.Fprintf(os.Stdout, "Response from `ScreenerApi.ListGridViewSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGridViewSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridViewSearchAdd** | [**GridViewSearchAdd**](GridViewSearchAdd.md) |  | 

### Return type

[**GridViewSearch**](GridViewSearch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrids

> Grids ListGrids(ctx).Include(include).GridsAdd(gridsAdd).Execute()





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
    include := "["info","score","analysis.raw_data","score.snowflake","analysis.extended.raw_data","analysis.extended.raw_data.insider_transactions","analysis.raw_data.insider_transactions"]" // string | string include (name or id) of the grids (optional)
    gridsAdd := *openapiclient.NewGridsAdd() // GridsAdd |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScreenerApi.ListGrids(context.Background()).Include(include).GridsAdd(gridsAdd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScreenerApi.ListGrids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGrids`: Grids
    fmt.Fprintf(os.Stdout, "Response from `ScreenerApi.ListGrids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGridsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | string include (name or id) of the grids | 
 **gridsAdd** | [**GridsAdd**](GridsAdd.md) |  | 

### Return type

[**Grids**](Grids.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

