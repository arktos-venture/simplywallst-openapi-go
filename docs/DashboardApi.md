# \DashboardApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndustryPerformance**](DashboardApi.md#ListIndustryPerformance) | **Get** /dashboard/industry-performance/{country} | 
[**ListInternationalMarket**](DashboardApi.md#ListInternationalMarket) | **Get** /dashboard/international-markets | 
[**ListMarketPerformance**](DashboardApi.md#ListMarketPerformance) | **Get** /dashboard/market-performance/{country} | 



## ListIndustryPerformance

> IndustryPerformance ListIndustryPerformance(ctx, country)



List properties of industryperformance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string**| string country (name or id) of the industryperformance | 

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

> InternationalMarket ListInternationalMarket(ctx, )



List properties of internationalmarket

### Required Parameters

This endpoint does not need any parameter.

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

> MarketPerformance ListMarketPerformance(ctx, country)



List properties of marketperformance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string**| string country (name or id) of the marketperformance | 

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

