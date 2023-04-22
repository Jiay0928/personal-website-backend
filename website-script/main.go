package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var projectList = []SideProject{
	{
		Title:      "Realtime Chatting Website",
		Tags:       []string{"web development", "react.js", "socket.io", "node.js"},
		Desc:       "build a chatroom allowing users to log in and sign up and create new chatrooms on the website",
		Date:       "Sept, 2022",
		KeyPoints:  []string{"Used React Framework for the frontend chat UI", "Used Node.js for backend and utilized socket.io library for messages transfer in real time"},
		VideoLink:  "https://s3.amazonaws.com/jiayingcui.com/videos/chatroom.mov",
		GithubLink: "https://github.com/Jiay0928/chat-backend-with-socket.io",
	},
	{
		Title: "Data Analysis Website",
		Tags:  []string{"web development", "react.js", "webpack", "javascript", "html", "css", "redux", "redux-thunk", "frontend"},
		Desc:  "Worked with teammates to create a data analysis web application which can connect to different databases",
		Date:  "Jul.- Ang., 2022",
		KeyPoints: []string{"Used react, redux and apache echarts to render tables and graphs of different datasets",
			"Worked with backend RESTful API to request and receive data using redux-thunk and async functions",
			"Used webpack to handle file loading and bundling",
		},
		VideoLink:  "https://s3.amazonaws.com/jiayingcui.com/videos/bividoe.mov",
		GithubLink: "https://github.com/Jiay0928/Bi_front",
	},
	{
		Title:      "Simple Tiktok GoLang Backend Project",
		Tags:       []string{"web development", "golang", "gorm", "backend", "mysql"},
		Date:       "May - Jun.,2022",
		Desc:       "Collaborated with teammates to create the backend for a simple app similar to TikTok using Golang",
		KeyPoints:  []string{"In charge of writing APIs for videos publishing and feeding using MVC structure", "Used gorm to interact with MySQL database"},
		VideoLink:  "",
		GithubLink: "https://github.com/tanghaijun666/SimpleTikTok.git",
	},
	{
		Title: "ConnectFour",
		Tags:  []string{"Kotlin", "MVC", "animation"},
		Date:  "Oct., 2022",
		Desc:  "built a  ConnectFour PC game",
		KeyPoints: []string{
			"Used JavaFx library and Animation to create the game", "implemented MVC structure",
		},
		VideoLink:  "https://s3.amazonaws.com/jiayingcui.com/videos/pc-game.mov",
		GithubLink: "https://github.com/Jiay0928/Connect-Four",
	},
	{
		Title: "Data Visualization PC App",
		Tags:  []string{"Kotlin", "MVC", "Data Visualization"},
		Date:  "Sept., 2022",
		Desc:  "create a data analysis computer software using kotlin",
		KeyPoints: []string{
			"able to modify data and render different graphs", "Used Kotlin built in graphing libaray to draw graph for the dataset", "binded data to graph using MVC structure"},
		VideoLink:  "https://s3.amazonaws.com/jiayingcui.com/videos/datagraph.mov",
		GithubLink: "https://github.com/Jiay0928/data-visualization-app",
	},
	{
		Title: "Notes - Android App",
		Tags:  append([]string{"Kotlin", "MVVM"}),
		Date:  "Oct.- Dec., 2022",
		Desc:  "Create a notes tracking Android app using Kotlin",
		KeyPoints: []string{
			"able to create, modify, archive notes in the application", "implemented usig MVVM structure",
		},
		VideoLink:  "https://s3.amazonaws.com/jiayingcui.com/videos/notes-phoneapp.mov",
		GithubLink: "https://github.com/Jiay0928/notes-android-app",
	},
}

type SideProject struct {
	Title      string   `json:"title" dynamodbav:"title"`
	Tags       []string `json:"Tags" dynamodbav:"Tags"`
	Desc       string   `json:"Desc" dynamodbav:"Desc"`
	Date       string   `json:"Date" dynamodbav:"Date"`
	KeyPoints  []string `json:"KeyPoints" dynamodbav:"KeyPoints"`
	VideoLink  string   `json:"VideoLink" dynamodbav:"VideoLink"`
	GithubLink string   `json:"GithubLink" dynamodbav:"GithubLink"`
}

func createItem(item SideProject) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("side_project"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

}

func main() {

	for _, value := range projectList {
		createItem(value)
	}

}
