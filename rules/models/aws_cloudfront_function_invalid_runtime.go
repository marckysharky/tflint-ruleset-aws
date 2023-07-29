// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudfrontFunctionInvalidRuntimeRule checks the pattern is valid
type AwsCloudfrontFunctionInvalidRuntimeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudfrontFunctionInvalidRuntimeRule returns new rule with default attributes
func NewAwsCloudfrontFunctionInvalidRuntimeRule() *AwsCloudfrontFunctionInvalidRuntimeRule {
	return &AwsCloudfrontFunctionInvalidRuntimeRule{
		resourceType:  "aws_cloudfront_function",
		attributeName: "runtime",
		enum: []string{
			"cloudfront-js-1.0",
			"cloudfront-js-2.0",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Name() string {
	return "aws_cloudfront_function_invalid_runtime"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudfrontFunctionInvalidRuntimeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as runtime`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
