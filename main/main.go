package main

import (
	"bufio"
	"fmt"
	"github.com/megaproaktiv/cdk2d2/monitor"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/exp/slog"
)

var CLI struct {
	Watch struct {
		// Force     bool `help:"Force removal."`
		Stack string   `arg:"" name:"stack" help:"Name of the stack." type:"string"`
		Paths []string `arg:"" name:"path" help:"Paths to cdk app." type:"path"`
	} `cmd:"" help:"Watch CDK generated CloudFormation resource status with d2 diagram. Generate output in <stackname>.d2. 2 second interval."`
	Generate struct {
		// Force     bool `help:"Force removal."`
		Stack string   `arg:"" name:"stack" help:"Name of the stack." type:"string"`
		Paths []string `arg:"" name:"path" help:"Paths to cdk app." type:"path"`
	} `cmd:"" help:"create generated CloudFormation resource status with d2 diagram.Output in <stackname>.d2 once."`
}

const cdkout = "cdk.out"
const manifestSubPath = cdkout + "/manifest.json"

//const templatePostfix = ".template.json"

func main() {
	var stackName *string
	var manifestPath *string
	oneTime := false
	//var templatePath *string
	ctx := kong.Parse(&CLI,
		kong.Name("cdk2d2"),
		kong.Description("Generate d2 diagramms with AWS-CDK."),
	)

	switch ctx.Command() {
	case "watch <stack> <path>":
		stackString := CLI.Watch.Stack
		stackName = &stackString
		manifestPathString := CLI.Watch.Paths[0]
		if !strings.HasSuffix(manifestPathString, "/") {
			manifestPathString += string(os.PathSeparator)
		}
		manifestPath = aws.String(manifestPathString + manifestSubPath)
		//templatePath = aws.String(manifestPathString +string(os.PathSeparator)+cdkout+string(os.PathSeparator)+stackString+templatePostfix)
	case "generate <stack> <path>":
		stackString := CLI.Generate.Stack
		stackName = &stackString
		manifestPathString := CLI.Generate.Paths[0]
		if !strings.HasSuffix(manifestPathString, "/") {
			manifestPathString += string(os.PathSeparator)
		}
		manifestPath = aws.String(manifestPathString + manifestSubPath)
		oneTime = true
		//templatePath = aws.String(manifestPathString +string(os.PathSeparator)+cdkout+string(os.PathSeparator)+stackString+templatePostfix)

	default:
		panic(ctx.Command())
	}

	manifest, err := monitor.ReadManifest(manifestPath, stackName)
	if err != nil {
		log.Println("Error reading manifest")
		os.Exit(1)
	}
	d2Name := *stackName + ".d2"
	slog.Info("Writing to:")
	slog.Info(d2Name)
	for {
		time.Sleep(4 * time.Second)
		file, err := os.Create(d2Name)
		if err != nil {
			fmt.Println("File does not exists or cannot be created")
			os.Exit(1)
		}
		stack, err := monitor.GetStatus(monitor.Client, stackName)
		if err != nil {
			log.Println("Error reading resources from stack, Wait 4 seconds.")
			file.Close()
			time.Sleep(4 * time.Second)
			continue
		}

		w := bufio.NewWriter(file)

		_, err = stack.Init(manifest)
		if err != nil {
			panic("Cant init graph")
		}
		_, err = stack.Graph(manifest, w)
		if err != nil {
			panic("Cant render graph")
		}

		resourcesInStack := 0
		readyResources := 0
		for _, r := range stack.Resources {
			resourcesInStack++
			if r.Status == monitor.StatusCreateComplete || r.Status == monitor.StatusUpdateComplete {
				readyResources++
			}
			if r.Visible {
				connections := manifest.Connections(r, stack)
				if len(connections) > 0 {
					for _, c := range connections {
						fmt.Fprintf(w, "%v \n", monitor.Connect(r.D2Id, c))
					}
				}
			}
		}
		fmt.Fprintf(w, "explanation: |md \n")
		fmt.Fprintf(w, " # %v\n", *stackName)
		percentReady := 100 * float64(readyResources) / float64(resourcesInStack)
		fmt.Fprintf(w, "- Resources: %d \n", int(resourcesInStack))
		fmt.Fprintf(w, "- Ready: %02d%% \n", int(percentReady))
		fmt.Fprintf(w, "| { near: top-left }\n")
		w.Flush()
		file.Close()
		if oneTime {
			break
		}
	}

}
