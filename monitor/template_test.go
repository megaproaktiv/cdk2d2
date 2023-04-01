package monitor_test

import (
	"testing"
	"github.com/megaproaktiv/cdk2d2/monitor"
	"github.com/aws/aws-sdk-go-v2/aws"
	"gotest.tools/assert"

)

func TestReadTemplate(t *testing.T) {

	path := aws.String("testdata/adot-xraystarter/adotstarter-auto.template.json")
	got, err := monitor.ReadTemplate(path)
	assert.NilError(t, err, "Template parsing should give not error.")
	assert.Equal(t, 30, len(*got))

}
