# EstimateCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **int64** |  | [optional] 
**Brokers** | Pointer to [**EstimateCoverageBrokers**](estimateCoverage_brokers.md) |  | [optional] 

## Methods

### NewEstimateCoverage

`func NewEstimateCoverage() *EstimateCoverage`

NewEstimateCoverage instantiates a new EstimateCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateCoverageWithDefaults

`func NewEstimateCoverageWithDefaults() *EstimateCoverage`

NewEstimateCoverageWithDefaults instantiates a new EstimateCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *EstimateCoverage) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EstimateCoverage) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EstimateCoverage) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EstimateCoverage) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetBrokers

`func (o *EstimateCoverage) GetBrokers() EstimateCoverageBrokers`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *EstimateCoverage) GetBrokersOk() (*EstimateCoverageBrokers, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *EstimateCoverage) SetBrokers(v EstimateCoverageBrokers)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *EstimateCoverage) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


