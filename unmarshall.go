package main

import (
    "fmt"
    "github.com/awslabs/goformation"
    "github.com/awslabs/goformation/cloudformation"
)

func main() {

    // Open a template from file (can be JSON or YAML)
    template, err := goformation.Open("template.yaml")

    // ...or provide one as a byte array ([]byte)
    template, err := goformation.ParseYAML(data)

    // You can then inspect all of the values
    for name, resource := range template.Resources {

        // E.g. Found a resource with name MyLambdaFunction and type AWS::Lambda::Function
        log.Printf("Found a resource with name %s and type %s", name, resource.Type)

    }

    // You can extract all resources of a certain type
    // Each AWS CloudFormation / SAM resource is a strongly typed struct
    functions := template.GetAllAWSLambdaFunctionResources()
    for name, function := range functions {

        // E.g. Found a AWS::Lambda::Function with name MyLambdaFunction and nodejs6.10 handler 
        log.Printf("Found a %s with name %s and %s handler", name, function.Type(), function.Handler)

    }

}
