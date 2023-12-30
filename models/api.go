package models

type (
	APIResponse[T any] struct {
		ID     string            `json:"id" example:"05aebb0d-0169-4bde-bc6d-8662e4680108"` // ID is the unique identifier of the request
		Method string            `json:"method" example:"link.list"`                        // Method is the text representation of the request (link.get)
		Params map[string]string `json:"params" example:"param:value,param2:value2"`        // Params is the parameters of the request
		Data   T                 `json:"data" swaggerignore:"true"`                         // Data is the data of the request
	}

	// ! APIResponseError is the response for errors.

	APIResponseError[I any] struct {
		ID     string            `json:"id" example:"05aebb0d-0169-4bde-bc6d-8662e4680108"` // ID is the unique identifier of the request
		Method string            `json:"method" example:"link.list"`                        // Method is the text representation of the request (link.get)
		Params map[string]string `json:"params" example:"param:value,param2:value2"`        // Params is the parameters of the request
		Error  I                 `json:"error" swaggerignore:"true"`                        // Error is the error of the request
	}

	APIResponseError500 struct {
		Code    int    `json:"code" example:"500"`
		Message string `json:"message" example:"Internal Server Error"`
	}

	APIResponseError404 struct {
		Code    int    `json:"code" example:"404"`
		Message string `json:"message" example:"Not Found"`
	}

	APIResponseError400 struct {
		Code    int    `json:"code" example:"400"`
		Message string `json:"message" example:"Bad Request"`
	}

	// 409 Conflict.
	APIResponseError409 struct {
		Code    int    `json:"code" example:"409"`
		Message string `json:"message" example:"Conflict"`
	}
)
