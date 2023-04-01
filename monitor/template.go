package monitor

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7"
	"golang.org/x/exp/slog"
)

// read cdk generated template

func ReadTemplate(path *string) (*cloudformation.Resources, error) {
	template, err := goformation.Open(*path)
	if err != nil {
		slog.Error("There was an error processing the template: %s", err)
		return nil, err
	}
	r := &template.Resources
	return r,nil
}
