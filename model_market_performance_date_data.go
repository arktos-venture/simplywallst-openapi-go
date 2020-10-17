/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MarketPerformanceDateData struct for MarketPerformanceDateData
type MarketPerformanceDateData struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TickerSymbol string `json:"ticker_symbol,omitempty"`
	UniqueSymbol string `json:"unique_symbol,omitempty"`
	CurrencyIso string `json:"currency_iso,omitempty"`
	LastSharePrice float32 `json:"last_share_price,omitempty"`
	Return7d int32 `json:"return_7d,omitempty"`
	Return30d int32 `json:"return_30d,omitempty"`
	Return90d int32 `json:"return_90d,omitempty"`
	Return1yrAbs int32 `json:"return_1yr_abs,omitempty"`
	Return3yrAbs int32 `json:"return_3yr_abs,omitempty"`
	Return5yrAbs int32 `json:"return_5yr_abs,omitempty"`
	CanonicalUrl string `json:"canonical_url,omitempty"`
}
