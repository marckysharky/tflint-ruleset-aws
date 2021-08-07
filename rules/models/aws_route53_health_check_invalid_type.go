// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53HealthCheckInvalidTypeRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53HealthCheckInvalidTypeRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidTypeRule() *AwsRoute53HealthCheckInvalidTypeRule {
	return &AwsRoute53HealthCheckInvalidTypeRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "type",
		enum: []string{
			"HTTP",
			"HTTPS",
			"HTTP_STR_MATCH",
			"HTTPS_STR_MATCH",
			"TCP",
			"CALCULATED",
			"CLOUDWATCH_METRIC",
			"RECOVERY_CONTROL",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidTypeRule) Name() string {
	return "aws_route53_health_check_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53HealthCheckInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
