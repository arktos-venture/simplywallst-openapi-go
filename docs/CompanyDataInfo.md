# CompanyDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WarningType** | Pointer to **int64** |  | [optional] 
**Industry** | Pointer to [**CompanyDataInfoIndustry**](companyDataInfo_industry.md) |  | [optional] 
**Fund** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Employees** | Pointer to **int32** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**YearFounded** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**CoverUrl** | Pointer to **string** |  | [optional] 
**CoverSmallUrl** | Pointer to **string** |  | [optional] 
**Ceo** | Pointer to [**CompanyDataInfoCeo**](companyDataInfo_ceo.md) |  | [optional] 
**LegalName** | Pointer to **string** |  | [optional] 

## Methods

### NewCompanyDataInfo

`func NewCompanyDataInfo() *CompanyDataInfo`

NewCompanyDataInfo instantiates a new CompanyDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDataInfoWithDefaults

`func NewCompanyDataInfoWithDefaults() *CompanyDataInfo`

NewCompanyDataInfoWithDefaults instantiates a new CompanyDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyDataInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyDataInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyDataInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyDataInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CompanyDataInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyDataInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyDataInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompanyDataInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWarningType

`func (o *CompanyDataInfo) GetWarningType() int64`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *CompanyDataInfo) GetWarningTypeOk() (*int64, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *CompanyDataInfo) SetWarningType(v int64)`

SetWarningType sets WarningType field to given value.

### HasWarningType

`func (o *CompanyDataInfo) HasWarningType() bool`

HasWarningType returns a boolean if a field has been set.

### GetIndustry

`func (o *CompanyDataInfo) GetIndustry() CompanyDataInfoIndustry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *CompanyDataInfo) GetIndustryOk() (*CompanyDataInfoIndustry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *CompanyDataInfo) SetIndustry(v CompanyDataInfoIndustry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *CompanyDataInfo) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetFund

`func (o *CompanyDataInfo) GetFund() bool`

GetFund returns the Fund field if non-nil, zero value otherwise.

### GetFundOk

`func (o *CompanyDataInfo) GetFundOk() (*bool, bool)`

GetFundOk returns a tuple with the Fund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFund

`func (o *CompanyDataInfo) SetFund(v bool)`

SetFund sets Fund field to given value.

### HasFund

`func (o *CompanyDataInfo) HasFund() bool`

HasFund returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyDataInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyDataInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyDataInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyDataInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *CompanyDataInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyDataInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyDataInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyDataInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyDataInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyDataInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyDataInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyDataInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmployees

`func (o *CompanyDataInfo) GetEmployees() int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *CompanyDataInfo) GetEmployeesOk() (*int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *CompanyDataInfo) SetEmployees(v int32)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *CompanyDataInfo) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetAddress

`func (o *CompanyDataInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyDataInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyDataInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CompanyDataInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetYearFounded

`func (o *CompanyDataInfo) GetYearFounded() int32`

GetYearFounded returns the YearFounded field if non-nil, zero value otherwise.

### GetYearFoundedOk

`func (o *CompanyDataInfo) GetYearFoundedOk() (*int32, bool)`

GetYearFoundedOk returns a tuple with the YearFounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearFounded

`func (o *CompanyDataInfo) SetYearFounded(v int32)`

SetYearFounded sets YearFounded field to given value.

### HasYearFounded

`func (o *CompanyDataInfo) HasYearFounded() bool`

HasYearFounded returns a boolean if a field has been set.

### GetUrl

`func (o *CompanyDataInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CompanyDataInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CompanyDataInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CompanyDataInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CompanyDataInfo) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CompanyDataInfo) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CompanyDataInfo) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CompanyDataInfo) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetCoverUrl

`func (o *CompanyDataInfo) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *CompanyDataInfo) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *CompanyDataInfo) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *CompanyDataInfo) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### GetCoverSmallUrl

`func (o *CompanyDataInfo) GetCoverSmallUrl() string`

GetCoverSmallUrl returns the CoverSmallUrl field if non-nil, zero value otherwise.

### GetCoverSmallUrlOk

`func (o *CompanyDataInfo) GetCoverSmallUrlOk() (*string, bool)`

GetCoverSmallUrlOk returns a tuple with the CoverSmallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverSmallUrl

`func (o *CompanyDataInfo) SetCoverSmallUrl(v string)`

SetCoverSmallUrl sets CoverSmallUrl field to given value.

### HasCoverSmallUrl

`func (o *CompanyDataInfo) HasCoverSmallUrl() bool`

HasCoverSmallUrl returns a boolean if a field has been set.

### GetCeo

`func (o *CompanyDataInfo) GetCeo() CompanyDataInfoCeo`

GetCeo returns the Ceo field if non-nil, zero value otherwise.

### GetCeoOk

`func (o *CompanyDataInfo) GetCeoOk() (*CompanyDataInfoCeo, bool)`

GetCeoOk returns a tuple with the Ceo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeo

`func (o *CompanyDataInfo) SetCeo(v CompanyDataInfoCeo)`

SetCeo sets Ceo field to given value.

### HasCeo

`func (o *CompanyDataInfo) HasCeo() bool`

HasCeo returns a boolean if a field has been set.

### GetLegalName

`func (o *CompanyDataInfo) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CompanyDataInfo) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CompanyDataInfo) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CompanyDataInfo) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


