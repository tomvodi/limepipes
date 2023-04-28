/*
 * Set and Tune API
 *
 * API for managing sets and tunes
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apimodel

type UpdateTune struct {
	Title string `json:"title" binding:"required"`

	Type string `json:"type,omitempty"`

	TimeSig string `json:"timeSig,omitempty"`

	Composer string `json:"composer,omitempty"`

	Arranger string `json:"arranger,omitempty"`
}

func (u UpdateTune) Validate() error {
	v := NewApimodelValidator()
	return v.Struct(u)
}
