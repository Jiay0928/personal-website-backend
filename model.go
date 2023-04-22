package main

// cui_side_project
import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const TableName = "side_project"

type SideProject struct {
	Title      string   `json:"title" dynamodbav:"title"`
	Tags       []string `json:"Tags" dynamodbav:"Tags"`
	Desc       string   `json:"Desc" dynamodbav:"Desc"`
	Date       string   `json:"Date" dynamodbav:"Date"`
	KeyPoints  []string `json:"KeyPoints" dynamodbav:"KeyPoints"`
	VideoLink  string   `json:"VideoLink" dynamodbav:"VideoLink"`
	GithubLink string   `json:"GithubLink" dynamodbav:"GithubLink"`
}

func getProjects(ctx context.Context) ([]SideProject, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
	if err != nil {
		log.Printf("fail at connection. Result: %#v", err)
		return nil, err
	}
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		TableName: aws.String("side_project"),
	}

	log.Printf("Calling Dynamodb with input: %v", input)
	result, err := svc.Scan(input)

	if err != nil {
		log.Printf("Executed GetItem DynamoDb successfully. Result: %#v", err)
		return nil, err
	}
	projectList := make([]SideProject, len(result.Items))
	for i, item := range result.Items {
		project := new(SideProject)
		err = dynamodbattribute.UnmarshalMap(item, project)
		if err != nil {
			return nil, err
		}
		projectList[i] = *project
	}

	return projectList, nil
}
