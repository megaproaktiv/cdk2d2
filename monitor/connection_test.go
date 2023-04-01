package monitor_test

import (
	"encoding/json"
	"fmt"
	"github.com/megaproaktiv/cdk2d2/monitor"
	"os"
	"testing"

	"github.com/megaproaktiv/awsmock"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"gotest.tools/assert"
)

func TestConnections(t *testing.T) {
	stackName := aws.String("adotstarter-auto")
	ListStackResourcesAdotFunction := func(ctx context.Context, params *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
		stackfile := "testdata/adot-xraystarter/list.json"
		data, err := os.ReadFile(stackfile)
		if err != nil {
			fmt.Println("File read error: ", err)
		}
		out := &cloudformation.ListStackResourcesOutput{}
		err = json.Unmarshal(data, out)
		if err != nil {
			fmt.Println("Unmarshel read error: ", err)
		}
		return out, nil
	}
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(ListStackResourcesAdotFunction)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())

	manifest, err := monitor.ReadManifest(aws.String("testdata/adot-xraystarter/manifest.json"), aws.String("adotstarter-auto"))
	assert.NilError(t, err)
	// assert.DeepEqual(t, Target, result)
	// Start with version
	assert.Equal(t, "21.0.0", manifest.Version)
	stack, err := monitor.GetStatus(client, stackName)
	assert.NilError(t, err)
	assert.Equal(t, 30, len(stack.SortedKeys))
	assert.Equal(t, true, monitor.Show(aws.String("items"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("adotstarter-ts"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("adotstarter-ts/ServiceRole"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("adotstarter-py"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("adotstarter-go"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("incoming"), manifest, stackName))
	assert.Equal(t, true, monitor.Show(aws.String("s3eventTopic"), manifest, stackName))

	lambdaPyResource := stack.Resources["adotstarterpy24CEA125"]
	lambdaPyConstructID := manifest.ConstructIdFromLogicalId(lambdaPyResource, stackName)
	assert.Equal(t, "adotstarter-py", lambdaPyConstructID)

	// Manifest entra:
	// "/adotstarter-auto/adotstarter-py": [
	// 	{
	// 	  "type": "Show",
	// 	  "data": "true"
	// 	},
	// 	{
	// 	  "type": "Connection",
	// 	  "data": "items"
	// 	}
	//   ],

	_, err = stack.Init(manifest)
	if err != nil {
		panic("Cant init graph")
	}
	connections := manifest.Connections(lambdaPyResource, stack)
	assert.Equal(t, "items07D08F4B", connections[0])

	lambdaTsResource := stack.Resources["adotstarterts5C66C01C"]
	connections = manifest.Connections(lambdaTsResource, stack)
	assert.Equal(t, "items07D08F4B", connections[0])

	lambdaGoResource := stack.Resources["adotstartergo2987B222"]
	connections = manifest.Connections(lambdaGoResource, stack)
	assert.Equal(t, "items07D08F4B", connections[0])

	incoming := stack.Resources["incoming0B397865"]
	connections = manifest.Connections(incoming, stack)
	assert.Equal(t, "s3eventTopicB05B6E9B", connections[0])

	s3eventTopic := stack.Resources["s3eventTopicB05B6E9B"]
	connections = manifest.Connections(s3eventTopic, stack)
	assert.Equal(t, "adotstarterts5C66C01C", connections[0])
	assert.Equal(t, "adotstartergo2987B222", connections[1])
	assert.Equal(t, "adotstarterpy24CEA125", connections[2])

}

func TestLogicalID(t *testing.T) {
	stackName := aws.String("adotstarter-auto")
	ListStackResourcesAdotFunction := func(ctx context.Context, params *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
		stackfile := "testdata/adot-xraystarter/list.json"
		data, err := os.ReadFile(stackfile)
		if err != nil {
			fmt.Println("File read error: ", err)
		}
		out := &cloudformation.ListStackResourcesOutput{}
		err = json.Unmarshal(data, out)
		if err != nil {
			fmt.Println("Unmarshel read error: ", err)
		}
		return out, nil
	}
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(ListStackResourcesAdotFunction)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())
	stack, err := monitor.GetStatus(client, aws.String("adotstarter-auto"))
	assert.NilError(t, err)

	manifest, err := monitor.ReadManifest(aws.String("testdata/adot-xraystarter/manifest.json"), aws.String("vpc"))
	assert.NilError(t, err)

	logicalIDMap := make(map[string]string, 0)
	d2IDMap := make(map[string]string, 0)
	// Enrich resources
	r := stack.Resources["adotstarterts5C66C01C"]
	constructID := manifest.ConstructIdFromLogicalId(r, stackName)
	assert.Equal(t, "adotstarter-ts", constructID)
	if len(constructID) > 0 {
		if monitor.Show(&constructID, manifest, stackName) {
			r.ConstructID = manifest.ConstructIdFromLogicalId(r, stackName)
			r.D2Id = r.LogicalResourceID
			d2IDMap[r.LogicalResourceID] = r.LogicalResourceID
			r.Visible = true
			logicalIDMap[r.ConstructID] = r.LogicalResourceID
		}
	}

	targetLogicalId := logicalIDMap["adotstarter-ts"]
	assert.Equal(t, "adotstarterts5C66C01C", targetLogicalId)

}
