package monitor

import (
	"fmt"
	"io"

	"oss.terrastruct.com/d2/d2graph"
)

// Write Graph
func (s *Stack) Graph(manifest *AssemblyManifest, w io.Writer ) (*d2graph.Graph, error){
	for _,stackResource := range s.SortedKeys {
		r := s.Resources[*stackResource]
		if len(r.ConstructID) > 0{
			if r.Visible {
				icon := Icon(&r.Type)
				d2shape := Transform(r, icon)
				fmt.Fprintf(w, "%v \n",*d2shape)
			}
		}
	}
	return nil, nil
}

func (s *Stack) Init(manifest *AssemblyManifest ) (*d2graph.Graph, error){
	for _,stackResource := range s.SortedKeys {
		r := s.Resources[*stackResource]
		constructID := manifest.ConstructIdFromLogicalId(r,s.Name)
		if len(constructID) > 0{
			if Show(&constructID, manifest,s.Name) {
				r.Visible = true
			}
			r.ConstructID = manifest.ConstructIdFromLogicalId(r,s.Name)
			// r.D2Id = manifest.D2ID(r, s.LogicalIDMap,s.Name)
			// s.D2IDMap[r.LogicalResourceID] = &r.D2Id
			lid := r.LogicalResourceID
			s.LogicalIDMap[r.ConstructID] = &lid
		}
	}
	// Container can only be created after all resources have logicalIds
	for _,stackResource := range s.SortedKeys {
		r := s.Resources[*stackResource]
		if len(r.ConstructID) > 0{
			r.D2Id = manifest.D2ID(r, s.LogicalIDMap,s.Name)
			s.D2IDMap[r.LogicalResourceID] = &r.D2Id
			lid := r.LogicalResourceID
			s.LogicalIDMap[r.ConstructID] = &lid
		}
	}
	return nil, nil
}

// Transform for the d2 output
func Transform(resource *CloudFormationResource, icon *string) *string {
	// shape := "rectangle"

	status := resource.Status
	var fill string

	switch status {
	case StatusCreateInProgress:
		fill = "yellow"
	case StatusCreateComplete:
		fill = "lightgreen"
	case StatusDeleteComplete:
		fill = "red"
	case StatusUpdateComplete:
		fill = "green"
	default:
		fill = "grey"
	}

	d2item := fmt.Sprintf("%v: %v{\n %v \n icon: %v \n style.fill:\"%v\" \n} \n",
		resource.D2Id,
		resource.ConstructID,
		resource.PhysicalResourceID,
		*icon,
		fill,
	)
	return &d2item
}
