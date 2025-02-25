/*
 * Sophos Public API
 *
 * Swagger Specifications for public APIs
 *
 * API version: beta
 * Contact: support@sophos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sophoscentral

type InstallerInfo struct {
	Beta        bool   `json:"beta,omitempty"`
	Command     string `json:"command,omitempty"`
	Platform    string `json:"platform,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Url         string `json:"url,omitempty"`
}
