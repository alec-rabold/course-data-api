package main

import (
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
	. "github.com/alec-rabold/EllucianBannerApi-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"
)

func handleRequest(event events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var data string
	var response events.APIGatewayProxyResponse
	queryParams := event.QueryStringParameters
	switch event.Path {
	case APIPathTerms:
		var requestModel request.TermsRequestModel
		mapstructure.Decode(queryParams, &requestModel)

	case APIPathSubjects:
		var requestModel request.SubjectsRequestModel
		mapstructure.Decode(queryParams, &requestModel)

	case APIPathCourses:
		var requestModel request.CoursesRequestModel
		mapstructure.Decode(queryParams, &requestModel)

	case APIPathSections:
		var requestModel request.SectionsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
	}
	response.Body = data

	return response
}

func main() {
	lambda.Start(handleRequest)
}
