package main

import (
	"fmt"

	"github.com/awslabs/goformation/cloudformation"
)

func main() {

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	// An an example SNS Topic
	template.Resources["MySNSTopic"] = &cloudformation.AWSSNSTopic{
		DisplayName: "test-sns-topic-display-name",
		TopicName:   "test-sns-topic-name",
		Subscription: []cloudformation.AWSSNSTopic_Subscription{
			cloudformation.AWSSNSTopic_Subscription{
				Endpoint: "test-sns-topic-subscription-endpoint",
				Protocol: "test-sns-topic-subscription-protocol",
			},
		},
	}

	// ...and a Route 53 Hosted Zone too
	template.Resources["MyRoute53HostedZone"] = &cloudformation.AWSRoute53HostedZone{
		Name: "gocloudformation.com",
	}

	// Let's see the JSON
	//j, err := template.JSON()
	//if err != nil {
	//	fmt.Printf("Failed to generate JSON: %s\n", err)
	//} else {
	//	fmt.Printf("%s\n", string(j))
	//}

	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}

}
