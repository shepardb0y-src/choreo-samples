/*
 * Todo API
 *
 * A simple Todo API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TodoRequest struct {

	Title string `json:"title"`

	Description string `json:"description,omitempty"`
}
