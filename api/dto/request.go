package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Request struct {
	Query string `validate:"required"`
	Pages int    `validate:"required"`
}

var (
	validate = validator.New()
)

func ValidateRequest(req Request) string {
	if err := validate.Struct(req); err != nil {
		return err.Error()
	}
	return ""
}

func RequestToString(req Request) string {
	return fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{
						"query_string": {
							"query": "%v"
						}
					}
				]
			}
		},
		"sort": [
			"-@timestamp"
		],
		"from": 0,
		"size": %v
	}
	`, req.Query, req.Pages)
}