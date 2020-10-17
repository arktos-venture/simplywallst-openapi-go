# \ScreenerApi

All URIs are relative to *https://api.simplywall.st*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGridViewSearch**](ScreenerApi.md#ListGridViewSearch) | **Post** /api/grid/view/search | 
[**ListGrids**](ScreenerApi.md#ListGrids) | **Post** /api/grid/filter | 



## ListGridViewSearch

> GridViewSearch ListGridViewSearch(ctx, optional)



List properties of gridviewsearch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGridViewSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGridViewSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridViewSearchAdd** | [**optional.Interface of GridViewSearchAdd**](GridViewSearchAdd.md)|  | 

### Return type

[**GridViewSearch**](gridViewSearch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrids

> Grids ListGrids(ctx, optional)



List properties of grids

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGridsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGridsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **optional.String**| string include (name or id) of the grids | 
 **gridsAdd** | [**optional.Interface of GridsAdd**](GridsAdd.md)|  | 

### Return type

[**Grids**](grids.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

