# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Iso2** | Pointer to **string** |  | [optional] 
**Iso3** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**CountryLinks**](country_links.md) |  | [optional] 
**Fields** | Pointer to [**CountryFields**](country_fields.md) |  | [optional] 

## Methods

### NewCountry

`func NewCountry() *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Country) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Country) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Country) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Country) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIso2

`func (o *Country) GetIso2() string`

GetIso2 returns the Iso2 field if non-nil, zero value otherwise.

### GetIso2Ok

`func (o *Country) GetIso2Ok() (*string, bool)`

GetIso2Ok returns a tuple with the Iso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2

`func (o *Country) SetIso2(v string)`

SetIso2 sets Iso2 field to given value.

### HasIso2

`func (o *Country) HasIso2() bool`

HasIso2 returns a boolean if a field has been set.

### GetIso3

`func (o *Country) GetIso3() string`

GetIso3 returns the Iso3 field if non-nil, zero value otherwise.

### GetIso3Ok

`func (o *Country) GetIso3Ok() (*string, bool)`

GetIso3Ok returns a tuple with the Iso3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3

`func (o *Country) SetIso3(v string)`

SetIso3 sets Iso3 field to given value.

### HasIso3

`func (o *Country) HasIso3() bool`

HasIso3 returns a boolean if a field has been set.

### GetType

`func (o *Country) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Country) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Country) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Country) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *Country) GetLinks() CountryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Country) GetLinksOk() (*CountryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Country) SetLinks(v CountryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Country) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetFields

`func (o *Country) GetFields() CountryFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Country) GetFieldsOk() (*CountryFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Country) SetFields(v CountryFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Country) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


