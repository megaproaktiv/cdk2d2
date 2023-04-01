package monitor_test

import (
  "github.com/megaproaktiv/cdk2d2/monitor"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/megaproaktiv/awsmock"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"gotest.tools/assert"
)

func TestContainer(t *testing.T) {
	stackName := aws.String("LoadBalancerStack")
	ListStackResourcesAdotFunction := func(ctx context.Context, params *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error){
		stackfile := "testdata/application-load-balancer/list.json"
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
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(ListStackResourcesAdotFunction)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())

	manifest, err := monitor.ReadManifest(aws.String("testdata/application-load-balancer/manifest.json"), stackName)
	assert.NilError(t,err)
	assert.Equal(t, "30.1.0", manifest.Version)
	stack, err := monitor.GetStatus(client, stackName)
	assert.NilError(t,err)
	assert.Equal(t, 36 , len(stack.SortedKeys))

	_ ,err = stack.Init(manifest)
	if err != nil{
		panic("Cant init graph")
	}

	currentD2ID := manifest.D2ID(stack.Resources["LB8A12904C"], stack.LogicalIDMap, stackName)
	assert.Equal(t, currentD2ID, "VPCB9E5F0B4.LB8A12904C")

}
