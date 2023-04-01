package monitor

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"golang.org/x/exp/slog"
)

var Client *cloudformation.Client
var programLevel = new(slog.LevelVar) // Info by default

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = cloudformation.NewFromConfig(cfg)

	// Loggin level
	h := slog.HandlerOptions{Level: programLevel}.NewTextHandler(os.Stderr)
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelInfo)
}

type Stack struct {
	Resources map[string]*CloudFormationResource
	SortedKeys []*string
	Name *string
	LogicalIDMap map[string]*string
	D2IDMap map[string]*string
}



type CloudFormationResource struct {
	ConstructID          string
	D2Id                 string
	Visible              bool
	LogicalResourceID    string
	PhysicalResourceID   string
	Status               string
	Type                 string
	Timestamp            time.Time
	ResourceStatusReason string
}

// StatusCreateComplete CloudFormation Status
const StatusCreateComplete = "CREATE_COMPLETE"

// StatusCreateInProgress CloudFormation Status
const StatusCreateInProgress = "CREATE_IN_PROGRESS"

// StatusDeleteComplete CloudFormation Status
const StatusDeleteComplete = "DELETE_COMPLETE"

// StatusUpdateComplete text for update
const StatusUpdateComplete = "UPDATE_COMPLETE"

const (
	// ColorDefault default color
	ColorDefault = "\x1b[39m"
	// ColorRed red for screen
	ColorRed = "\x1b[91m"
	// ColorGreen green for screen
	ColorGreen = "\x1b[32m"
	// ColorBlue blue for screen
	ColorBlue = "\x1b[94m"
	// ColorGray for screen
	ColorGray = "\x1b[90m"
)

// Get State of all resources
func GetStatus(client *cloudformation.Client, name *string) (*Stack, error) {
	data := make(map[string]*CloudFormationResource)
	params := &cloudformation.ListStackResourcesInput{
		StackName: name,
	}
	myStack := &Stack{
		Resources: data,
	}
	resp, err := client.ListStackResources(context.TODO(), params)
	if err != nil {
		fmt.Printf("Stacks: Stack %v is not readable.\n", *name)
		return nil, err
	}
	sortedLogicalIds := make([]*string, 0, len(resp.StackResourceSummaries))
	
	counter := 0
	for _, resource := range resp.StackResourceSummaries {
		item := &CloudFormationResource{
			LogicalResourceID:  *resource.LogicalResourceId,
			PhysicalResourceID: "",
			Status:             string(resource.ResourceStatus),
			Type:               *resource.ResourceType,
		}
		data[*resource.LogicalResourceId] = item
		sortedLogicalIds =append(sortedLogicalIds,  resource.LogicalResourceId)
		counter++
	}
	myStack.SortedKeys = sortedLogicalIds
	myStack.Name = name
	myStack.D2IDMap = make(map[string]*string)
	myStack.LogicalIDMap = make(map[string]*string)
	return myStack,  nil
}

