# Companies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Company**](Company.md) |  | [optional] 
**Meta** | Pointer to [**CompanyMeta**](CompanyMeta.md) |  | [optional] 

## Methods

### NewCompanies

`func NewCompanies() *Companies`

NewCompanies instantiates a new Companies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompaniesWithDefaults

`func NewCompaniesWithDefaults() *Companies`

NewCompaniesWithDefaults instantiates a new Companies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Companies) GetData() Company`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Companies) GetDataOk() (*Company, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Companies) SetData(v Company)`

SetData sets Data field to given value.

### HasData

`func (o *Companies) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *Companies) GetMeta() CompanyMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Companies) GetMetaOk() (*CompanyMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Companies) SetMeta(v CompanyMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Companies) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


