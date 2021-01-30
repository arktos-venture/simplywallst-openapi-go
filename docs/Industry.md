# Industry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **int64** |  | [optional] 

## Methods

### NewIndustry

`func NewIndustry() *Industry`

NewIndustry instantiates a new Industry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndustryWithDefaults

`func NewIndustryWithDefaults() *Industry`

NewIndustryWithDefaults instantiates a new Industry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Industry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Industry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Industry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Industry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Industry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Industry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Industry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Industry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParent

`func (o *Industry) GetParent() int64`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Industry) GetParentOk() (*int64, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Industry) SetParent(v int64)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Industry) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


