/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GridViewSearchAdd struct for GridViewSearchAdd
type GridViewSearchAdd struct {
	Id int32 `json:"id,omitempty"`
	NoResultIfLimit bool `json:"no_result_if_limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
	Size int32 `json:"size,omitempty"`
	State string `json:"state,omitempty"`
	Rules string `json:"rules,omitempty"`
}
