/* 
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Swaggeroauth2TokenParameters struct {

	// in: formData
	ClientId string `json:"client_id,omitempty"`

	// in: formData
	Code string `json:"code,omitempty"`

	// in: formData
	GrantType string `json:"grant_type"`

	// in: formData
	RedirectUri string `json:"redirect_uri,omitempty"`
}
