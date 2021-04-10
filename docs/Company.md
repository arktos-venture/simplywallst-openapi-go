# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**TradingItemId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to [**OneOfstringinteger**](oneOf&lt;string,integer&gt;.md) |  | [optional] 
**Slug** | Pointer to [**OneOfstringinteger**](oneOf&lt;string,integer&gt;.md) |  | [optional] 
**ExchangeSymbol** | Pointer to **string** |  | [optional] 
**UniqueSymbol** | Pointer to **string** |  | [optional] 
**PrimaryTicker** | Pointer to **bool** |  | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**CanonicalUrl** | Pointer to **string** |  | [optional] 
**PrimaryCanonicalUrl** | Pointer to **NullableString** |  | [optional] 
**IsSearchable** | Pointer to **bool** |  | [optional] 
**IsinSymbol** | Pointer to **string** |  | [optional] 
**Analysis** | Pointer to [**CompanyAnalysis**](CompanyAnalysis.md) |  | [optional] 
**Info** | Pointer to [**CompanyInfo**](CompanyInfo.md) |  | [optional] 
**Score** | Pointer to [**CompanyScore**](CompanyScore.md) |  | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Company) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Company) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Company) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Company) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetTradingItemId

`func (o *Company) GetTradingItemId() int64`

GetTradingItemId returns the TradingItemId field if non-nil, zero value otherwise.

### GetTradingItemIdOk

`func (o *Company) GetTradingItemIdOk() (*int64, bool)`

GetTradingItemIdOk returns a tuple with the TradingItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingItemId

`func (o *Company) SetTradingItemId(v int64)`

SetTradingItemId sets TradingItemId field to given value.

### HasTradingItemId

`func (o *Company) HasTradingItemId() bool`

HasTradingItemId returns a boolean if a field has been set.

### GetName

`func (o *Company) GetName() OneOfstringinteger`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*OneOfstringinteger, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v OneOfstringinteger)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *Company) GetSlug() OneOfstringinteger`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Company) GetSlugOk() (*OneOfstringinteger, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Company) SetSlug(v OneOfstringinteger)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Company) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetExchangeSymbol

`func (o *Company) GetExchangeSymbol() string`

GetExchangeSymbol returns the ExchangeSymbol field if non-nil, zero value otherwise.

### GetExchangeSymbolOk

`func (o *Company) GetExchangeSymbolOk() (*string, bool)`

GetExchangeSymbolOk returns a tuple with the ExchangeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeSymbol

`func (o *Company) SetExchangeSymbol(v string)`

SetExchangeSymbol sets ExchangeSymbol field to given value.

### HasExchangeSymbol

`func (o *Company) HasExchangeSymbol() bool`

HasExchangeSymbol returns a boolean if a field has been set.

### GetUniqueSymbol

`func (o *Company) GetUniqueSymbol() string`

GetUniqueSymbol returns the UniqueSymbol field if non-nil, zero value otherwise.

### GetUniqueSymbolOk

`func (o *Company) GetUniqueSymbolOk() (*string, bool)`

GetUniqueSymbolOk returns a tuple with the UniqueSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSymbol

`func (o *Company) SetUniqueSymbol(v string)`

SetUniqueSymbol sets UniqueSymbol field to given value.

### HasUniqueSymbol

`func (o *Company) HasUniqueSymbol() bool`

HasUniqueSymbol returns a boolean if a field has been set.

### GetPrimaryTicker

`func (o *Company) GetPrimaryTicker() bool`

GetPrimaryTicker returns the PrimaryTicker field if non-nil, zero value otherwise.

### GetPrimaryTickerOk

`func (o *Company) GetPrimaryTickerOk() (*bool, bool)`

GetPrimaryTickerOk returns a tuple with the PrimaryTicker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTicker

`func (o *Company) SetPrimaryTicker(v bool)`

SetPrimaryTicker sets PrimaryTicker field to given value.

### HasPrimaryTicker

`func (o *Company) HasPrimaryTicker() bool`

HasPrimaryTicker returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Company) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Company) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Company) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Company) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCanonicalUrl

`func (o *Company) GetCanonicalUrl() string`

GetCanonicalUrl returns the CanonicalUrl field if non-nil, zero value otherwise.

### GetCanonicalUrlOk

`func (o *Company) GetCanonicalUrlOk() (*string, bool)`

GetCanonicalUrlOk returns a tuple with the CanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUrl

`func (o *Company) SetCanonicalUrl(v string)`

SetCanonicalUrl sets CanonicalUrl field to given value.

### HasCanonicalUrl

`func (o *Company) HasCanonicalUrl() bool`

HasCanonicalUrl returns a boolean if a field has been set.

### GetPrimaryCanonicalUrl

`func (o *Company) GetPrimaryCanonicalUrl() string`

GetPrimaryCanonicalUrl returns the PrimaryCanonicalUrl field if non-nil, zero value otherwise.

### GetPrimaryCanonicalUrlOk

`func (o *Company) GetPrimaryCanonicalUrlOk() (*string, bool)`

GetPrimaryCanonicalUrlOk returns a tuple with the PrimaryCanonicalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCanonicalUrl

`func (o *Company) SetPrimaryCanonicalUrl(v string)`

SetPrimaryCanonicalUrl sets PrimaryCanonicalUrl field to given value.

### HasPrimaryCanonicalUrl

`func (o *Company) HasPrimaryCanonicalUrl() bool`

HasPrimaryCanonicalUrl returns a boolean if a field has been set.

### SetPrimaryCanonicalUrlNil

`func (o *Company) SetPrimaryCanonicalUrlNil(b bool)`

 SetPrimaryCanonicalUrlNil sets the value for PrimaryCanonicalUrl to be an explicit nil

### UnsetPrimaryCanonicalUrl
`func (o *Company) UnsetPrimaryCanonicalUrl()`

UnsetPrimaryCanonicalUrl ensures that no value is present for PrimaryCanonicalUrl, not even an explicit nil
### GetIsSearchable

`func (o *Company) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *Company) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *Company) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *Company) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### GetIsinSymbol

`func (o *Company) GetIsinSymbol() string`

GetIsinSymbol returns the IsinSymbol field if non-nil, zero value otherwise.

### GetIsinSymbolOk

`func (o *Company) GetIsinSymbolOk() (*string, bool)`

GetIsinSymbolOk returns a tuple with the IsinSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsinSymbol

`func (o *Company) SetIsinSymbol(v string)`

SetIsinSymbol sets IsinSymbol field to given value.

### HasIsinSymbol

`func (o *Company) HasIsinSymbol() bool`

HasIsinSymbol returns a boolean if a field has been set.

### GetAnalysis

`func (o *Company) GetAnalysis() CompanyAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *Company) GetAnalysisOk() (*CompanyAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *Company) SetAnalysis(v CompanyAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *Company) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetInfo

`func (o *Company) GetInfo() CompanyInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Company) GetInfoOk() (*CompanyInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Company) SetInfo(v CompanyInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Company) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetScore

`func (o *Company) GetScore() CompanyScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Company) GetScoreOk() (*CompanyScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Company) SetScore(v CompanyScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *Company) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


