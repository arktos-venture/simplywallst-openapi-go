# Grids

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Company**](Company.md) |  | [optional] 
**Meta** | Pointer to [**GridsMeta**](grids_meta.md) |  | [optional] 

## Methods

### NewGrids

`func NewGrids() *Grids`

NewGrids instantiates a new Grids object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridsWithDefaults

`func NewGridsWithDefaults() *Grids`

NewGridsWithDefaults instantiates a new Grids object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Grids) GetData() []Company`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Grids) GetDataOk() (*[]Company, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Grids) SetData(v []Company)`

SetData sets Data field to given value.

### HasData

`func (o *Grids) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *Grids) GetMeta() GridsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Grids) GetMetaOk() (*GridsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Grids) SetMeta(v GridsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Grids) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


