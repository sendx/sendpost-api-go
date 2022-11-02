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

type ModelsBackOffConfiguration struct {
	ConcurrentConnections int64 `json:"concurrentConnections,omitempty"`
	ConcurrentConnectionsType *ModelsBackOffDecreaseType `json:"concurrentConnectionsType,omitempty"`
	SendPerDay int64 `json:"sendPerDay,omitempty"`
	SendPerDayType *ModelsBackOffDecreaseType `json:"sendPerDayType,omitempty"`
	SendPerHour int64 `json:"sendPerHour,omitempty"`
	SendPerHourType *ModelsBackOffDecreaseType `json:"sendPerHourType,omitempty"`
	SendPerMinute int64 `json:"sendPerMinute,omitempty"`
	SendPerMinuteType *ModelsBackOffDecreaseType `json:"sendPerMinuteType,omitempty"`
}
