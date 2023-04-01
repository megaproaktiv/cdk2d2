package monitor

import (
	"encoding/json"
	"os"

	"golang.org/x/exp/slog"
)

type ConstructId string
type LogicalResourceId string





// read cdk manifest
func ReadManifest(filename *string, stackname *string) (*AssemblyManifest, error){
	content, err := os.ReadFile(*filename)
    if err != nil {
        slog.Error("Error when opening file: ", err)
    }
 
    var manifest = AssemblyManifest{
    	Version:   "",
    	Artifacts: make(map[string]*ArtifactManifest),
    }
	
    err = json.Unmarshal(content, &manifest)
    if err != nil {
        slog.Error("Error during Unmarshal(): ", err)
    }
	// Assembly can contain many stacks
	//manifest.Stack = *stackname
	return &manifest,nil
}
