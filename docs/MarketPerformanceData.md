# MarketPerformanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Top** | Pointer to [**map[string]MarketPerformanceDate**](marketPerformanceDate.md) |  | [optional] 
**Worst** | Pointer to [**map[string]MarketPerformanceDate**](marketPerformanceDate.md) |  | [optional] 

## Methods

### NewMarketPerformanceData

`func NewMarketPerformanceData() *MarketPerformanceData`

NewMarketPerformanceData instantiates a new MarketPerformanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPerformanceDataWithDefaults

`func NewMarketPerformanceDataWithDefaults() *MarketPerformanceData`

NewMarketPerformanceDataWithDefaults instantiates a new MarketPerformanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTop

`func (o *MarketPerformanceData) GetTop() map[string]MarketPerformanceDate`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *MarketPerformanceData) GetTopOk() (*map[string]MarketPerformanceDate, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *MarketPerformanceData) SetTop(v map[string]MarketPerformanceDate)`

SetTop sets Top field to given value.

### HasTop

`func (o *MarketPerformanceData) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetWorst

`func (o *MarketPerformanceData) GetWorst() map[string]MarketPerformanceDate`

GetWorst returns the Worst field if non-nil, zero value otherwise.

### GetWorstOk

`func (o *MarketPerformanceData) GetWorstOk() (*map[string]MarketPerformanceDate, bool)`

GetWorstOk returns a tuple with the Worst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorst

`func (o *MarketPerformanceData) SetWorst(v map[string]MarketPerformanceDate)`

SetWorst sets Worst field to given value.

### HasWorst

`func (o *MarketPerformanceData) HasWorst() bool`

HasWorst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


