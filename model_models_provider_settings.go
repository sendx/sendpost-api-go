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

type ModelsProviderSettings struct {
	BackOffConfiguration *ModelsBackOffConfiguration `json:"backOffConfiguration,omitempty"`
	BackOffTrigger *ModelsBackOffTrigger `json:"backOffTrigger,omitempty"`
	MaxConcurrentConnections int64 `json:"maxConcurrentConnections,omitempty"`
	MaxSendPerDay int64 `json:"maxSendPerDay,omitempty"`
	MaxSendPerHour int64 `json:"maxSendPerHour,omitempty"`
	MaxSendPerMinute int64 `json:"maxSendPerMinute,omitempty"`
	Name string `json:"name,omitempty"`
}
