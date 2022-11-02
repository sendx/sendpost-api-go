/*
 * SendPost API
 *
 * Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.
 *
 * API version: 1.0.0
 * Contact: hello@sendpost.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsDetailedAlert struct {
	Active bool `json:"active,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	EntityValue string `json:"entityValue,omitempty"`
	Id int64 `json:"id,omitempty"`
	Label *ModelsAlertLabel `json:"label,omitempty"`
	Name string `json:"name,omitempty"`
	Parameter int64 `json:"parameter,omitempty"`
	Provider string `json:"provider,omitempty"`
	Threshold string `json:"threshold,omitempty"`
}