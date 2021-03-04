# IndustryTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**map[string]map[string]IndustryValues**](map.md) |  | [optional] 

## Methods

### NewIndustryTree

`func NewIndustryTree() *IndustryTree`

NewIndustryTree instantiates a new IndustryTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndustryTreeWithDefaults

`func NewIndustryTreeWithDefaults() *IndustryTree`

NewIndustryTreeWithDefaults instantiates a new IndustryTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IndustryTree) GetData() map[string]map[string]IndustryValues`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IndustryTree) GetDataOk() (*map[string]map[string]IndustryValues, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IndustryTree) SetData(v map[string]map[string]IndustryValues)`

SetData sets Data field to given value.

### HasData

`func (o *IndustryTree) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


