/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CompanyAnalysisExtendedRawDataDividendData struct for CompanyAnalysisExtendedRawDataDividendData
type CompanyAnalysisExtendedRawDataDividendData struct {
	PayDate                 *int64   `json:"pay_date,omitempty"`
	RecordDate              *int64   `json:"record_date,omitempty"`
	NextDate                *int64   `json:"next_date,omitempty"`
	Amount                  *float64 `json:"amount,omitempty"`
	ReportingCurrencyAmount *float64 `json:"reporting_currency_amount,omitempty"`
	SplitAdjustedAmount     *float64 `json:"split_adjusted_amount,omitempty"`
	Date                    *int64   `json:"date,omitempty"`
	CurrencyIso             *string  `json:"currency_iso,omitempty"`
}

// NewCompanyAnalysisExtendedRawDataDividendData instantiates a new CompanyAnalysisExtendedRawDataDividendData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAnalysisExtendedRawDataDividendData() *CompanyAnalysisExtendedRawDataDividendData {
	this := CompanyAnalysisExtendedRawDataDividendData{}
	return &this
}

// NewCompanyAnalysisExtendedRawDataDividendDataWithDefaults instantiates a new CompanyAnalysisExtendedRawDataDividendData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAnalysisExtendedRawDataDividendDataWithDefaults() *CompanyAnalysisExtendedRawDataDividendData {
	this := CompanyAnalysisExtendedRawDataDividendData{}
	return &this
}

// GetPayDate returns the PayDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetPayDate() int64 {
	if o == nil || o.PayDate == nil {
		var ret int64
		return ret
	}
	return *o.PayDate
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetPayDateOk() (*int64, bool) {
	if o == nil || o.PayDate == nil {
		return nil, false
	}
	return o.PayDate, true
}

// HasPayDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasPayDate() bool {
	if o != nil && o.PayDate != nil {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given int64 and assigns it to the PayDate field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetPayDate(v int64) {
	o.PayDate = &v
}

// GetRecordDate returns the RecordDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetRecordDate() int64 {
	if o == nil || o.RecordDate == nil {
		var ret int64
		return ret
	}
	return *o.RecordDate
}

// GetRecordDateOk returns a tuple with the RecordDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetRecordDateOk() (*int64, bool) {
	if o == nil || o.RecordDate == nil {
		return nil, false
	}
	return o.RecordDate, true
}

// HasRecordDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasRecordDate() bool {
	if o != nil && o.RecordDate != nil {
		return true
	}

	return false
}

// SetRecordDate gets a reference to the given int64 and assigns it to the RecordDate field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetRecordDate(v int64) {
	o.RecordDate = &v
}

// GetNextDate returns the NextDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetNextDate() int64 {
	if o == nil || o.NextDate == nil {
		var ret int64
		return ret
	}
	return *o.NextDate
}

// GetNextDateOk returns a tuple with the NextDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetNextDateOk() (*int64, bool) {
	if o == nil || o.NextDate == nil {
		return nil, false
	}
	return o.NextDate, true
}

// HasNextDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasNextDate() bool {
	if o != nil && o.NextDate != nil {
		return true
	}

	return false
}

// SetNextDate gets a reference to the given int64 and assigns it to the NextDate field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetNextDate(v int64) {
	o.NextDate = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetAmount() float64 {
	if o == nil || o.Amount == nil {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetAmountOk() (*float64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetAmount(v float64) {
	o.Amount = &v
}

// GetReportingCurrencyAmount returns the ReportingCurrencyAmount field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetReportingCurrencyAmount() float64 {
	if o == nil || o.ReportingCurrencyAmount == nil {
		var ret float64
		return ret
	}
	return *o.ReportingCurrencyAmount
}

// GetReportingCurrencyAmountOk returns a tuple with the ReportingCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetReportingCurrencyAmountOk() (*float64, bool) {
	if o == nil || o.ReportingCurrencyAmount == nil {
		return nil, false
	}
	return o.ReportingCurrencyAmount, true
}

// HasReportingCurrencyAmount returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasReportingCurrencyAmount() bool {
	if o != nil && o.ReportingCurrencyAmount != nil {
		return true
	}

	return false
}

// SetReportingCurrencyAmount gets a reference to the given float64 and assigns it to the ReportingCurrencyAmount field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetReportingCurrencyAmount(v float64) {
	o.ReportingCurrencyAmount = &v
}

// GetSplitAdjustedAmount returns the SplitAdjustedAmount field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetSplitAdjustedAmount() float64 {
	if o == nil || o.SplitAdjustedAmount == nil {
		var ret float64
		return ret
	}
	return *o.SplitAdjustedAmount
}

// GetSplitAdjustedAmountOk returns a tuple with the SplitAdjustedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetSplitAdjustedAmountOk() (*float64, bool) {
	if o == nil || o.SplitAdjustedAmount == nil {
		return nil, false
	}
	return o.SplitAdjustedAmount, true
}

// HasSplitAdjustedAmount returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasSplitAdjustedAmount() bool {
	if o != nil && o.SplitAdjustedAmount != nil {
		return true
	}

	return false
}

// SetSplitAdjustedAmount gets a reference to the given float64 and assigns it to the SplitAdjustedAmount field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetSplitAdjustedAmount(v float64) {
	o.SplitAdjustedAmount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetDate() int64 {
	if o == nil || o.Date == nil {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetDateOk() (*int64, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetDate(v int64) {
	o.Date = &v
}

// GetCurrencyIso returns the CurrencyIso field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetCurrencyIso() string {
	if o == nil || o.CurrencyIso == nil {
		var ret string
		return ret
	}
	return *o.CurrencyIso
}

// GetCurrencyIsoOk returns a tuple with the CurrencyIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) GetCurrencyIsoOk() (*string, bool) {
	if o == nil || o.CurrencyIso == nil {
		return nil, false
	}
	return o.CurrencyIso, true
}

// HasCurrencyIso returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataDividendData) HasCurrencyIso() bool {
	if o != nil && o.CurrencyIso != nil {
		return true
	}

	return false
}

// SetCurrencyIso gets a reference to the given string and assigns it to the CurrencyIso field.
func (o *CompanyAnalysisExtendedRawDataDividendData) SetCurrencyIso(v string) {
	o.CurrencyIso = &v
}

func (o CompanyAnalysisExtendedRawDataDividendData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayDate != nil {
		toSerialize["pay_date"] = o.PayDate
	}
	if o.RecordDate != nil {
		toSerialize["record_date"] = o.RecordDate
	}
	if o.NextDate != nil {
		toSerialize["next_date"] = o.NextDate
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.ReportingCurrencyAmount != nil {
		toSerialize["reporting_currency_amount"] = o.ReportingCurrencyAmount
	}
	if o.SplitAdjustedAmount != nil {
		toSerialize["split_adjusted_amount"] = o.SplitAdjustedAmount
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.CurrencyIso != nil {
		toSerialize["currency_iso"] = o.CurrencyIso
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyAnalysisExtendedRawDataDividendData struct {
	value *CompanyAnalysisExtendedRawDataDividendData
	isSet bool
}

func (v NullableCompanyAnalysisExtendedRawDataDividendData) Get() *CompanyAnalysisExtendedRawDataDividendData {
	return v.value
}

func (v *NullableCompanyAnalysisExtendedRawDataDividendData) Set(val *CompanyAnalysisExtendedRawDataDividendData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAnalysisExtendedRawDataDividendData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAnalysisExtendedRawDataDividendData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAnalysisExtendedRawDataDividendData(val *CompanyAnalysisExtendedRawDataDividendData) *NullableCompanyAnalysisExtendedRawDataDividendData {
	return &NullableCompanyAnalysisExtendedRawDataDividendData{value: val, isSet: true}
}

func (v NullableCompanyAnalysisExtendedRawDataDividendData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAnalysisExtendedRawDataDividendData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
