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

type ModelsAccount struct {
	AlertSlackEndpoint string `json:"alertSlackEndpoint,omitempty"`
	ApiKey string `json:"apiKey,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Created int64 `json:"created,omitempty"`
	CurrentEmailServiceProvider string `json:"currentEmailServiceProvider,omitempty"`
	Id int64 `json:"id,omitempty"`
	IncidentSlackEndpoint string `json:"incidentSlackEndpoint,omitempty"`
	Industry string `json:"industry,omitempty"`
	IsCanceled bool `json:"isCanceled,omitempty"`
	IsLastPaymentFailed bool `json:"isLastPaymentFailed,omitempty"`
	IsUpgraded bool `json:"isUpgraded,omitempty"`
	LockThreshold int64 `json:"lockThreshold,omitempty"`
	Locked bool `json:"locked,omitempty"`
	LogoURL string `json:"logoURL,omitempty"`
	Name string `json:"name,omitempty"`
	OnboardCFinished bool `json:"onboardCFinished,omitempty"`
	OnboardQAnswered bool `json:"onboardQAnswered,omitempty"`
	SendingVolumePerMonth string `json:"sendingVolumePerMonth,omitempty"`
	SlackToken string `json:"slackToken,omitempty"`
	StripeBasePriceId string `json:"stripeBasePriceId,omitempty"`
	StripeUsagePriceId string `json:"stripeUsagePriceId,omitempty"`
}
