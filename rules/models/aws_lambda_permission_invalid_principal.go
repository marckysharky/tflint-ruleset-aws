// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaPermissionInvalidPrincipalRule checks the pattern is valid
type AwsLambdaPermissionInvalidPrincipalRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidPrincipalRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidPrincipalRule() *AwsLambdaPermissionInvalidPrincipalRule {
	return &AwsLambdaPermissionInvalidPrincipalRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "principal",
		pattern:       regexp.MustCompile(`^[^\s]+$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidPrincipalRule) Name() string {
	return "aws_lambda_permission_invalid_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidPrincipalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidPrincipalRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidPrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidPrincipalRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\s]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
