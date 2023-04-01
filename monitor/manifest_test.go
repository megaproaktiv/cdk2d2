package monitor_test

import (
  "github.com/megaproaktiv/cdk2d2/monitor"
	"testing"
	"gotest.tools/assert"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestReadManifest(t *testing.T) {

	treeTarget := &monitor.ArtifactManifest{
		Type:        "cdk:tree",
		Environment: "",
	}

	result, err := monitor.ReadManifest(aws.String("testdata/manifest-1.json"), aws.String("xraystarter"))
	assert.NilError(t,err)
	// assert.DeepEqual(t, Target, result)
	// Start with version
	t.Log("Version")
	assert.Equal(t, "21.0.0", result.Version)
	t.Log("Tree")
	assert.Equal(t, treeTarget.Type, result.Artifacts["Tree"].Type)
	xraystarter := result.Artifacts["xraystarter"]
	metadata := xraystarter.Metadata["/xraystarter/xraystarter-ts"]
	assert.Equal(t,metadata[0].Data.(bool), true)
	assert.Equal(t,metadata[1].Data.(string), "LambdaNameTS")

}

func TestContructID(t *testing.T) {

	stack := aws.String("adotstarter-auto")

	result, err := monitor.ReadManifest(aws.String("testdata/adot-xraystarter/manifest.json"), aws.String("adotstarter-auto"))
	assert.NilError(t,err)

	//Check Bucket resource:
	// "/adotstarter-auto/incoming/Resource": [
	// 	{
	// 	  "type": "aws:cdk:logicalId",
	// 	  "data": "incoming0B397865"
	// 	}
	//   ],
	t.Logf("Test counstructID <incoming>\n")
	constructIDwant := "incoming"
	target := &monitor.CloudFormationResource{
		LogicalResourceID: "incoming0B397865",
	}
	constructIDgot := result.ConstructIdFromLogicalId(target,stack)
	assert.Equal(t, constructIDwant, constructIDgot)


}
