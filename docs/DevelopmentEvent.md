# DevelopmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to [**DevelopmentEventTypes**](DevelopmentEventTypes.md) |  | [optional] 

## Methods

### NewDevelopmentEvent

`func NewDevelopmentEvent() *DevelopmentEvent`

NewDevelopmentEvent instantiates a new DevelopmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevelopmentEventWithDefaults

`func NewDevelopmentEventWithDefaults() *DevelopmentEvent`

NewDevelopmentEventWithDefaults instantiates a new DevelopmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *DevelopmentEvent) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DevelopmentEvent) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DevelopmentEvent) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *DevelopmentEvent) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetId

`func (o *DevelopmentEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevelopmentEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevelopmentEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DevelopmentEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *DevelopmentEvent) GetPriority() bool`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DevelopmentEvent) GetPriorityOk() (*bool, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DevelopmentEvent) SetPriority(v bool)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DevelopmentEvent) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTypes

`func (o *DevelopmentEvent) GetTypes() DevelopmentEventTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DevelopmentEvent) GetTypesOk() (*DevelopmentEventTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DevelopmentEvent) SetTypes(v DevelopmentEventTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DevelopmentEvent) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


