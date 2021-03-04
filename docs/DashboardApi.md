# \DashboardApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndustryPerformance**](DashboardApi.md#ListIndustryPerformance) | **Get** /dashboard/industry-performance/{country} | 
[**ListInternationalMarket**](DashboardApi.md#ListInternationalMarket) | **Get** /dashboard/international-markets | 
[**ListMarketPerformance**](DashboardApi.md#ListMarketPerformance) | **Get** /dashboard/market-performance/{country} | 



## ListIndustryPerformance

> IndustryPerformance ListIndustryPerformance(ctx, country).Execute()





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
    country := "["us","hk"]" // string | string country (name or id) of the industryperformance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.ListIndustryPerformance(context.Background(), country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.ListIndustryPerformance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndustryPerformance`: IndustryPerformance
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.ListIndustryPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string** | string country (name or id) of the industryperformance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIndustryPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndustryPerformance**](industryPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternationalMarket

> InternationalMarket ListInternationalMarket(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.ListInternationalMarket(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.ListInternationalMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInternationalMarket`: InternationalMarket
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.ListInternationalMarket`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInternationalMarketRequest struct via the builder pattern


### Return type

[**InternationalMarket**](internationalMarket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketPerformance

> MarketPerformance ListMarketPerformance(ctx, country).Execute()





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
    country := "["us","hk"]" // string | string country (name or id) of the marketperformance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardApi.ListMarketPerformance(context.Background(), country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.ListMarketPerformance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMarketPerformance`: MarketPerformance
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.ListMarketPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string** | string country (name or id) of the marketperformance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMarketPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketPerformance**](marketPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

