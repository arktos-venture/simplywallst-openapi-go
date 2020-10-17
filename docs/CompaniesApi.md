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

> DevelopmentInfo ListDevelopmentInfo(ctx, include, optional)



List properties of developmentinfo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**include** | **string**| string include (name or id) of the developmentinfo | 
 **optional** | ***ListDevelopmentInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevelopmentInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **developmentInfoAdd** | [**optional.Interface of DevelopmentInfoAdd**](DevelopmentInfoAdd.md)|  | 

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

> Developments ListDevelopments(ctx, id)



List properties of developments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the developments | 

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

> EstimateCoverages ListEstimateCoverages(ctx, id, optional)



List properties of estimatecoverages

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the estimatecoverages | 
 **optional** | ***ListEstimateCoveragesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEstimateCoveragesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| string version (name or id) of the estimatecoverages | 

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

> News ListNews(ctx, id, optional)



List properties of news

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the news | 
 **optional** | ***ListNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNewsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simplywallSt** | **optional.Bool**| boolean simplywall.st (name or id) of the news | 
 **version** | **optional.String**| string version (name or id) of the news | 

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

> Ownerships ListOwnerships(ctx, id)



List properties of ownerships

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the ownerships | 

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

> Prices ListPrices(ctx, id, optional)



List properties of prices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the prices | 
 **optional** | ***ListPricesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPricesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTimestamp** | **optional.Int32**| integer start_timestamp (name or id) of the prices | 

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

> Companies ReadCompanies(ctx, exchange, sector, ticker, company, optional)



Read properties of companies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string**| string exchange (name or id) of the companies | 
**sector** | **string**| string sector (name or id) of the companies | 
**ticker** | **string**| string ticker (name or id) of the companies | 
**company** | **string**| string company (name or id) of the companies | 
 **optional** | ***ReadCompaniesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCompaniesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **include** | **optional.String**| string include (name or id) of the companies | 
 **version** | **optional.String**| string version (name or id) of the companies | 

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

