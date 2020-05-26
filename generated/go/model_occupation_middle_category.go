/*
 * job_offer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OccupationMiddleCategory struct {

	// 職業大分類コード
	IscoCode string `json:"iscoCode,omitempty"`

	// 職業中分類
	IsmcoCode string `json:"ismcoCode,omitempty"`

	// 職業中分類名
	IsmcoName string `json:"ismcoName,omitempty"`
}
