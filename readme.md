# AWS lambda function perform as a backend to query dynamodb
To compile the files 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go lambda.go model.go