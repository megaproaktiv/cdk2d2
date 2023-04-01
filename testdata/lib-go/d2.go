package d2

import (
		"github.com/aws/constructs-go/constructs/v10"
		"github.com/aws/aws-sdk-go-v2/aws"
)

type NodeAble interface {
	Node() constructs.Node
}	

func Show( construct NodeAble){
	construct.Node().AddMetadata(aws.String("Show"), aws.String("true"),nil)
}

func Connection( construct NodeAble, target NodeAble){
	construct.Node().AddMetadata(aws.String("Connection"), target.Node().Id(),nil)
}