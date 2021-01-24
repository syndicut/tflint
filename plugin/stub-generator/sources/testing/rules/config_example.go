package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ConfigExampleRule checks whether ...
type ConfigExampleRule struct{}

// NewConfigExampleRule returns a new rule
func NewConfigExampleRule() *ConfigExampleRule {
	return &ConfigExampleRule{}
}

// Name returns the rule name
func (r *ConfigExampleRule) Name() string {
	return "config_example"
}

// Enabled returns whether the rule is enabled by default
func (r *ConfigExampleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ConfigExampleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ConfigExampleRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ConfigExampleRule) Check(runner tflint.Runner) error {
	_, err := runner.Config()
	if err != nil {
		return err
	}
	return nil
}
