# InternationalMarketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Top** | Pointer to [**map[string]InternationalMarketReturn**](internationalMarketReturn.md) |  | [optional] 
**Worst** | Pointer to [**map[string]InternationalMarketReturn**](internationalMarketReturn.md) |  | [optional] 

## Methods

### NewInternationalMarketData

`func NewInternationalMarketData() *InternationalMarketData`

NewInternationalMarketData instantiates a new InternationalMarketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalMarketDataWithDefaults

`func NewInternationalMarketDataWithDefaults() *InternationalMarketData`

NewInternationalMarketDataWithDefaults instantiates a new InternationalMarketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTop

`func (o *InternationalMarketData) GetTop() map[string]InternationalMarketReturn`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *InternationalMarketData) GetTopOk() (*map[string]InternationalMarketReturn, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *InternationalMarketData) SetTop(v map[string]InternationalMarketReturn)`

SetTop sets Top field to given value.

### HasTop

`func (o *InternationalMarketData) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetWorst

`func (o *InternationalMarketData) GetWorst() map[string]InternationalMarketReturn`

GetWorst returns the Worst field if non-nil, zero value otherwise.

### GetWorstOk

`func (o *InternationalMarketData) GetWorstOk() (*map[string]InternationalMarketReturn, bool)`

GetWorstOk returns a tuple with the Worst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorst

`func (o *InternationalMarketData) SetWorst(v map[string]InternationalMarketReturn)`

SetWorst sets Worst field to given value.

### HasWorst

`func (o *InternationalMarketData) HasWorst() bool`

HasWorst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


