// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmDevicePoolInvalidNameRule checks the pattern is valid
type AwsDevicefarmDevicePoolInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsDevicefarmDevicePoolInvalidNameRule returns new rule with default attributes
func NewAwsDevicefarmDevicePoolInvalidNameRule() *AwsDevicefarmDevicePoolInvalidNameRule {
	return &AwsDevicefarmDevicePoolInvalidNameRule{
		resourceType:  "aws_devicefarm_device_pool",
		attributeName: "name",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsDevicefarmDevicePoolInvalidNameRule) Name() string {
	return "aws_devicefarm_device_pool_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmDevicePoolInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmDevicePoolInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmDevicePoolInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmDevicePoolInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
