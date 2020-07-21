package main

import (
	"context"
	"encoding/json"

	"github.com/alec-rabold/EllucianBannerApi-go/pkg/client"
	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/request"
	. "github.com/alec-rabold/EllucianBannerApi-go/pkg/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data []byte
	queryParams := event.QueryStringParameters
	client, err := client.NewAPIClient(client.EllucianBanner)
	if err != nil {
		log.Errorf("error instantiating API client, err: %v", err)
	}
	switch event.Path {
	case APIPathColleges:
		data, _ = json.Marshal(client.GetColleges())
	case APIPathTerms:
		var requestModel request.TermsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(client.GetTerms(requestModel))
	case APIPathSubjects:
		var requestModel request.SubjectsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(client.GetSubjects(requestModel))
	case APIPathCourses:
		var requestModel request.CoursesRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(client.GetCourses(requestModel))
	case APIPathSections:
		var requestModel request.SectionsRequestModel
		mapstructure.Decode(queryParams, &requestModel)
		data, _ = json.Marshal(client.GetSections(requestModel))
	}
	return events.APIGatewayProxyResponse{Body: string(data), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
