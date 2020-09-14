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

type ModelsEipPool struct {
	Ips []ModelsEip `json:"ips,omitempty"`
	Name string `json:"name,omitempty"`
	RoutingMapping *interface{} `json:"routingMapping,omitempty"`
	RoutingStrategy int64 `json:"routingStrategy,omitempty"`
}