package monitor

import (
	"fmt"
	"strings"
	"golang.org/x/exp/slog"
)

// Show whether the resource is displayed
func Show(constructID *string, resources *AssemblyManifest, stackname *string) bool {
	var stackArtifacts = resources.Artifacts[*stackname]
	var metadata = stackArtifacts.Metadata

	for key, value := range metadata {
		foundShow := false
		var metaDataEntry MetadataEntry
		if len(value) > 0 {
			for _, item := range value {
				metaDataEntry = *item
				if metaDataEntry.Type == "Show" {
					showValue := metaDataEntry.Data.(string)
					if strings.EqualFold(showValue, "true") {
						foundShow = true
					}
				}
			}
		}
		if foundShow {
			manifestConstructId := ""
			path := strings.Split(key, "/")
			// /$stack/$cid
			// "/vpc/baseVPC"
			if len(path) == 3 {	
				manifestConstructId = path[2]
			}else{
				// "/vpc/baseVPC/privatewebaSubnet1" 
				manifestConstructId = path[2] + "/" + path[3]				
				
			}
			
			if manifestConstructId == *constructID {		
				return true
			}
		}
	}

	return false
}

// Relationships

func (manifest *AssemblyManifest) Connections(r *CloudFormationResource,  stack *Stack) []string {
	connections := make([]string, 0)
	var stackArtifacts = manifest.Artifacts[*stack.Name]
	if stackArtifacts == nil {
		return connections
	}
	var metadata = stackArtifacts.Metadata
	constructID := manifest.ConstructIdFromLogicalId(r, stack.Name)

	for key, value := range metadata {
		path := strings.Split(key, "/")
		// "/adotstarter-auto/incoming"
		if len(path) == 3 &&
			path[1] == *stack.Name &&
			path[2] == constructID {

			var metaDataEntry MetadataEntry

			if len(value) > 0 {
				for _, item := range value {
					metaDataEntry = *item
					if metaDataEntry.Type == "Connection" {
						slog.Debug("Start connections")
						targetConstructID := metaDataEntry.Data.(string)
						slog.Debug("cid: ", targetConstructID)
						targetLogicalId := stack.LogicalIDMap[targetConstructID]
						if targetLogicalId != nil &&  len(*targetLogicalId) > 0 {
							targetLogicalId := *stack.LogicalIDMap[targetConstructID]
							targetD2Id := stack.D2IDMap[targetLogicalId]
							// if target is not shown
							if len(*targetD2Id) > 0 {
								connections = append(connections, *targetD2Id)
							}
						}
					}
				}
			} else {
				continue
			}
		}

		//   "/vpc/baseVPC/privatewebaSubnet1"
		if len(path) == 4 &&
			path[1] == *stack.Name &&
			path[3] == constructID {
			var metaDataEntry MetadataEntry

			if len(value) > 0 {
				for _, item := range value {
					metaDataEntry = *item
					if metaDataEntry.Type == "Connection" {
						targetConstructID := metaDataEntry.Data.(string)
						targetLogicalId := stack.LogicalIDMap[targetConstructID]
						targetD2Id := stack.D2IDMap[*targetLogicalId]
						// if target is not shown
						if len(*targetD2Id) > 0 {
							connections = append(connections, *targetD2Id)
						}
					}
				}
			} else {
				continue
			}

		}

	}

	return connections

}

func Connect(fromD2Id string, toD2Id string) string {
	var con string
	if len(fromD2Id) > 0 && len(toD2Id) > 0 {
		con = fmt.Sprintf("\n %v -> %v \n", fromD2Id, toD2Id)
	} else {
		con = ""
	}
	return con
}

func (r *AssemblyManifest) ConstructIdFromLogicalId(cr *CloudFormationResource, stackname *string) string {

	logicalId := cr.LogicalResourceID
	var stackArtifacts = r.Artifacts[*stackname]
	if stackArtifacts.Metadata == nil {
		return ""
	}
	var metadata = stackArtifacts.Metadata
	
	for key, value := range metadata {
		foundLogicalId := false
		// foundpath := false

		// An Entry with "type": "aws:cdk:logicalId",
		var metaDataEntry MetadataEntry
		if len(value) > 0 {
			for _, item := range value {
				metaDataEntry = *item
				if metaDataEntry.Type == "aws:cdk:logicalId" {
					logicalIdMetadata := metaDataEntry.Data.(string)
					if logicalIdMetadata == logicalId {
						//return constructId
						foundLogicalId = true
					}
				}
			}
		}

		if foundLogicalId {
			// Resource
			path := strings.Split(key, "/")

			if( len(path) == 3 ){
				constructId := path[2]
				return constructId
			}

			if len(path) == 4 && strings.HasSuffix(key, "Resource") {
				constructId := path[2]
				return constructId
			}
			if len(path) == 4 && (cr.Type == "AWS::AutoScaling::AutoScalingGroup" && strings.HasSuffix(key, "ASG")) {
				constructId := path[2]
				return constructId
			}
			// "/vpc/baseVPC/privatewebaSubnet1/Subnet" => baseVPC/privatewebaSubnet1
			if len(path) == 5 && (strings.HasSuffix(key, "Subnet") && (cr.Type == "AWS::EC2::Subnet")) {
				constructIdString := path[2] + "/" + path[3]
				constructId := constructIdString
				lowerCid := strings.ToLower(constructId)
				if strings.Contains(lowerCid, "priv") {
					cr.Type = "AWS::EC2::Subnet::Private"
				}
				if strings.Contains(lowerCid, "pub") {
					cr.Type = "AWS::EC2::Subnet::Public"
				}
				return constructId
			}
		}
	}
	return ""
}

// Compose d2 id from metadata entries
// "/cloudair/monolithSG": [
//
//		{
//		  "type": "Show",
//		  "data": "true"
//		},
//		{
//		  "type": "Container",
//		  "data": "MonolithVPC"
//		}
//	  ],
func (m *AssemblyManifest) D2ID(r *CloudFormationResource, logicalIDMap map[string]*string, stack *string) string {
	constructId := r.ConstructID
	stackArtifacts := m.Artifacts[*stack]

	metadata := stackArtifacts.Metadata
	for key, value := range metadata {
		path := strings.Split(key, "/")
		// "/cloudair/monolithSG"
		if len(path) == 3 &&
			path[1] == *stack &&
			path[2] == constructId {
			var metaDataEntry MetadataEntry

			if len(value) > 0 {
				for _, item := range value {
					metaDataEntry = *item
					if metaDataEntry.Type == "Container" {
						containerConstructId := metaDataEntry.Data.(string)
						containerLogicalId := logicalIDMap[containerConstructId]
						fqD2Id := fmt.Sprintf("%v.%v", *containerLogicalId, r.LogicalResourceID)

						return fqD2Id
					}
				}
			}
		} else {
			continue
		}
	}

	return r.LogicalResourceID
}
