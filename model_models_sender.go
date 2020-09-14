/*
 * SendPost API
 *
 * SendPost API to send transactional emails reliably
 *
 * API version: 1.0.0
 * Contact: hello@sendx.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsSender struct {
	Created int64 `json:"created,omitempty"`
	Domain string `json:"domain,omitempty"`
	FromEmail string `json:"fromEmail,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ReplyToEmail string `json:"replyToEmail,omitempty"`
	Verified bool `json:"verified,omitempty"`
}
