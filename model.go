package main

// cui_side_project
import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const TableName = "side_project"

type SideProject struct {
	Id         int      `json:"id" dynamodbav:"id"`
	Title      string   `json:"title" dynamodbav:"title"`
	Tags       []string `json:"Tags" dynamodbav:"Tags"`
	Desc       string   `json:"Desc" dynamodbav:"Desc"`
	Date       string   `json:"Date" dynamodbav:"Date"`
	KeyPoints  []string `json:"KeyPoints" dynamodbav:"KeyPoints"`
	VideoLink  string   `json:"VideoLink" dynamodbav:"VideoLink"`
	GithubLink string   `json:"GithubLink" dynamodbav:"GithubLink"`
}

var db dynamodb.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db = *dynamodb.NewFromConfig(sdkConfig)
}

func getProjects(ctx context.Context) ([]SideProject, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("side_project"),
	}

	log.Printf("Calling Dynamodb with input: %v", input)
	result, err := db.Scan(ctx, input)
	if err != nil {
		log.Printf("Executed GetItem DynamoDb successfully. Result: %#v", err)
		return nil, err
	}
	projectList := make([]SideProject, len(result.Items))
	for i, item := range result.Items {
		project := new(SideProject)
		err = attributevalue.UnmarshalMap(item, project)
		if err != nil {
			return nil, err
		}
		projectList[i] = *project
	}

	return projectList, nil
}
