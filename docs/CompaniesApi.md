# \CompaniesApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDevelopmentInfo**](CompaniesApi.md#ListDevelopmentInfo) | **Post** /api/company/developments/info | 
[**ListDevelopments**](CompaniesApi.md#ListDevelopments) | **Get** /api/company/developments/{id} | 
[**ListEstimateCoverages**](CompaniesApi.md#ListEstimateCoverages) | **Get** /api/company/estimates/coverage/{id} | 
[**ListNews**](CompaniesApi.md#ListNews) | **Get** /api/news/{id} | 
[**ListOwnerships**](CompaniesApi.md#ListOwnerships) | **Get** /api/company/ownership/shareholders/{id} | 
[**ListPrices**](CompaniesApi.md#ListPrices) | **Get** /api/company/price/{id} | 
[**ReadCompanies**](CompaniesApi.md#ReadCompanies) | **Get** /api/company/stocks/{exchange}/{sector}/{ticker}/{company} | 



## ListDevelopmentInfo

> DevelopmentInfo ListDevelopmentInfo(ctx).Include(include).DevelopmentInfoAdd(developmentInfoAdd).Execute()





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
    include := "["info","score","analysis.raw_data","score.snowflake","analysis.extended.raw_data","analysis.extended.raw_data.insider_transactions","analysis.raw_data.insider_transactions"]" // string | string include (name or id) of the developmentinfo
    developmentInfoAdd := *openapiclient.NewDevelopmentInfoAdd() // DevelopmentInfoAdd |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListDevelopmentInfo(context.Background()).Include(include).DevelopmentInfoAdd(developmentInfoAdd).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListDevelopmentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevelopmentInfo`: DevelopmentInfo
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListDevelopmentInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevelopmentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | string include (name or id) of the developmentinfo | 
 **developmentInfoAdd** | [**DevelopmentInfoAdd**](DevelopmentInfoAdd.md) |  | 

### Return type

[**DevelopmentInfo**](developmentInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevelopments

> Developments ListDevelopments(ctx, id).Execute()





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
    id := TODO // string | string id (name or id) of the developments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListDevelopments(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListDevelopments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevelopments`: Developments
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListDevelopments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | string id (name or id) of the developments | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevelopmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Developments**](developments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEstimateCoverages

> EstimateCoverages ListEstimateCoverages(ctx, id).Version(version).Execute()





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
    id := TODO // string | string id (name or id) of the estimatecoverages
    version := "[2]" // string | string version (name or id) of the estimatecoverages (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListEstimateCoverages(context.Background(), id).Version(version).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListEstimateCoverages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEstimateCoverages`: EstimateCoverages
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListEstimateCoverages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | string id (name or id) of the estimatecoverages | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEstimateCoveragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | string version (name or id) of the estimatecoverages | 

### Return type

[**EstimateCoverages**](estimateCoverages.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNews

> News ListNews(ctx, id).SimplywallSt(simplywallSt).Version(version).Execute()





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
    id := TODO // string | string id (name or id) of the news
    simplywallSt := false // bool | boolean simplywall.st (name or id) of the news (optional)
    version := "[2]" // string | string version (name or id) of the news (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListNews(context.Background(), id).SimplywallSt(simplywallSt).Version(version).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListNews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNews`: News
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListNews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | string id (name or id) of the news | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simplywallSt** | **bool** | boolean simplywall.st (name or id) of the news | 
 **version** | **string** | string version (name or id) of the news | 

### Return type

[**News**](news.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOwnerships

> Ownerships ListOwnerships(ctx, id).Execute()





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
    id := TODO // string | string id (name or id) of the ownerships

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListOwnerships(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListOwnerships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOwnerships`: Ownerships
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListOwnerships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | string id (name or id) of the ownerships | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOwnershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ownerships**](ownerships.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrices

> Prices ListPrices(ctx, id).StartTimestamp(startTimestamp).Execute()





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
    id := TODO // string | string id (name or id) of the prices
    startTimestamp := int64(789) // int64 | integer start_timestamp (name or id) of the prices (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListPrices(context.Background(), id).StartTimestamp(startTimestamp).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrices`: Prices
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | string id (name or id) of the prices | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTimestamp** | **int64** | integer start_timestamp (name or id) of the prices | 

### Return type

[**Prices**](prices.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCompanies

> Companies ReadCompanies(ctx, exchange, sector, ticker, company).Include(include).Version(version).Execute()





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
    exchange := "["us","hk"]" // string | string exchange (name or id) of the companies
    sector := "["energy"]" // string | string sector (name or id) of the companies
    ticker := "["nyse-tot"]" // string | string ticker (name or id) of the companies
    company := "["total","apple"]" // string | string company (name or id) of the companies
    include := "["info","score","analysis.raw_data","score.snowflake","analysis.extended.raw_data","analysis.extended.raw_data.insider_transactions","analysis.raw_data.insider_transactions"]" // string | string include (name or id) of the companies (optional)
    version := "[2]" // string | string version (name or id) of the companies (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ReadCompanies(context.Background(), exchange, sector, ticker, company).Include(include).Version(version).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ReadCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCompanies`: Companies
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ReadCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string** | string exchange (name or id) of the companies | 
**sector** | **string** | string sector (name or id) of the companies | 
**ticker** | **string** | string ticker (name or id) of the companies | 
**company** | **string** | string company (name or id) of the companies | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **include** | **string** | string include (name or id) of the companies | 
 **version** | **string** | string version (name or id) of the companies | 

### Return type

[**Companies**](companies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

