# \CompetitorsApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCompetitors**](CompetitorsApi.md#ListCompetitors) | **Get** /api/competitors/{id} | 



## ListCompetitors

> Competitors ListCompetitors(ctx, id, optional)



List properties of competitors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| string id (name or id) of the competitors | 
 **optional** | ***ListCompetitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCompetitorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| string include (name or id) of the competitors | 
 **version** | **optional.String**| string version (name or id) of the competitors | 

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

