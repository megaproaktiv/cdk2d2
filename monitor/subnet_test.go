package monitor_test

import (
  "github.com/megaproaktiv/cdk2d2/monitor"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/megaproaktiv/awsmock"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"context"
	"gotest.tools/assert"
)

var ListStackResourcesVpcFunction = func(ctx context.Context, params *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error){
	stackfile := "testdata/list-vpc.json"
	data, err := os.ReadFile(stackfile)
	if err != nil {
		fmt.Println("File read error: ", err)
	}
	out := &cloudformation.ListStackResourcesOutput{}
	err = json.Unmarshal(data, out);
	if err != nil {
		fmt.Println("Unmarshel read error: ", err)
	}
	return out, nil
}

func TestVpcVpc(t *testing.T) {
	stackName := aws.String("vpc")

	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(ListStackResourcesVpcFunction)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())

	manifest, err := monitor.ReadManifest(aws.String("testdata/manifest-vpc.json"), aws.String("vpc"))
	assert.NilError(t,err)
	// assert.DeepEqual(t, Target, result)
	// Start with version
	assert.Equal(t, "13.0.0", manifest.Version)


	stack, err := monitor.GetStatus(client, aws.String("vpc"))
	assert.NilError(t,err)

	assert.Equal(t, len(stack.Resources), 78)
	assert.Equal(t, true, monitor.Show(aws.String("baseVPC"), manifest,stackName))
	assert.Equal(t, false, monitor.Show(aws.String("ssmVpcId"), manifest,stackName))

	vpcResource := stack.Resources["baseVPC9645629F"]
	vpcConstructID := manifest.ConstructIdFromLogicalId(vpcResource,stackName)
	assert.Equal(t,"baseVPC", vpcConstructID)


}


func TestVpcSubnet(t *testing.T) {

	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(ListStackResourcesVpcFunction)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())

	manifest, err := monitor.ReadManifest(aws.String("testdata/manifest-vpc.json"), aws.String("vpc"))
	assert.NilError(t,err)
	assert.Equal(t, "13.0.0", manifest.Version)

	stackName := aws.String("vpc")

	// MockStackStatus
	stack, err := monitor.GetStatus(client, stackName)
	assert.NilError(t,err)

	assert.Equal(t, len(stack.Resources), 78)

	// Manifest:
	// "/vpc/baseVPC/privatewebaSubnet1": [
	// 	{
	// 	  "type": "Show",
	// 	  "data": "true"
	// 	}
	//   ],
	subnetShowResource := stack.Resources["baseVPCprivatewebaSubnet1Subnet694F3021"]
	subnetConstructId := manifest.ConstructIdFromLogicalId(subnetShowResource,stackName)

	assert.Equal(t, "baseVPC/privatewebaSubnet1",subnetConstructId)
	assert.Equal(t, true, monitor.Show(aws.String("baseVPC/privatewebaSubnet1"), manifest,stackName))
}
