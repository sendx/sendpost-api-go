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

type ModelsAccountWebhook struct {
	Clicked bool `json:"clicked,omitempty"`
	Created int64 `json:"created,omitempty"`
	Delivered bool `json:"delivered,omitempty"`
	Dropped bool `json:"dropped,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	HardBounced bool `json:"hardBounced,omitempty"`
	Id int64 `json:"id,omitempty"`
	Opened bool `json:"opened,omitempty"`
	Processed bool `json:"processed,omitempty"`
	SoftBounced bool `json:"softBounced,omitempty"`
	Spam bool `json:"spam,omitempty"`
	Unsubscribed bool `json:"unsubscribed,omitempty"`
	Url string `json:"url,omitempty"`
}