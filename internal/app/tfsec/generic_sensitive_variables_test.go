package tfsec

import (
	"testing"

	"github.com/liamg/tfsec/internal/app/tfsec/checks"
)

func Test_AWSSensitiveVariables(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode checks.Code
		mustExcludeResultCode checks.Code
	}{
		{
			name: "check sensitive variable with value",
			source: `
variable "db_password" {
	default = "something"
}`,
			mustIncludeResultCode: checks.GenericSensitiveVariables,
		},
		{
			name: "check sensitive variable without value",
			source: `
variable "db_password" {
	default = ""
}`,
			mustExcludeResultCode: checks.GenericSensitiveVariables,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
