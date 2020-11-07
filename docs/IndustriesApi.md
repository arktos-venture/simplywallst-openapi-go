# \IndustriesApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndustry**](IndustriesApi.md#ListIndustry) | **Get** /api/industry/company/{id} | 
[**ListIndustryCountry**](IndustriesApi.md#ListIndustryCountry) | **Get** /industry/0/{country} | 
[**ListIndustryTree**](IndustriesApi.md#ListIndustryTree) | **Get** /api/industry/tree | 



## ListIndustry

> Industry ListIndustry(ctx, id)



List properties of industry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| string id (name or id) of the industry | 

### Return type

[**Industry**](industry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndustryCountry

> IndustryCountry ListIndustryCountry(ctx, country)



List properties of industrycountry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string**| string country (name or id) of the industrycountry | 

### Return type

[**IndustryCountry**](industryCountry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndustryTree

> IndustryTree ListIndustryTree(ctx, )



List properties of industrytree

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**IndustryTree**](industryTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

