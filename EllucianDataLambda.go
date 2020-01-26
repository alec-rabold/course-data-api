package main

import (
	"context"
	"encoding/json"
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
	"github.com/alec-rabold/EllucianBannerApi-go/service"
	. "github.com/alec-rabold/EllucianBannerApi-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"
)

func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data []byte
	queryParams := event.QueryStringParameters
	service, _ := service.NewAPIService()
	switch event.Path {
	case APIPathColleges:
		data, _ = json.Marshal(service.GetColleges())
	case APIPathTerms:
		var requestModel request.TermsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(service.GetTerms(requestModel))
	case APIPathSubjects:
		var requestModel request.SubjectsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(service.GetSubjects(requestModel))
	case APIPathCourses:
		var requestModel request.CoursesRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(service.GetCourses(requestModel))
	case APIPathSections:
		var requestModel request.SectionsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(service.GetSections(requestModel))
	}
	return events.APIGatewayProxyResponse{Body: string(data), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
