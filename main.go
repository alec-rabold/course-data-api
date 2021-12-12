package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	apiv1 "github.com/alec-rabold/course-data-api/pkg/api/v1"
	"github.com/alec-rabold/course-data-api/pkg/client/ellucian"
	"github.com/alec-rabold/course-data-api/pkg/crawler"
	model "github.com/alec-rabold/course-data-api/pkg/model/request"
	"github.com/alec-rabold/course-data-api/pkg/router"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log, ctx, err := initalizeLogger(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to instantiate logger: %v", err)
	}
	log.Info("incoming request", "request", request)

	r := router.NewRouter(apiv1.BasePath)

	return r.Handle(ctx, request)
}

// TODO: the parameters handling can probably be built into the router
func newRouter(ctx context.Context) {
	r := router.NewRouter(apiv1.BasePath)

	// Currently only Ellucian Banner is supported so default to it here
	client := ellucian.NewClient(crawler.NewCrawler(
		// TODO: may need to set a phony user agent
		crawler.WithContentType("application/x-www-form-urlencoded"),
	))

	r.Route(http.MethodGet, apiv1.PathColleges, func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return marshalResponse(client.GetColleges(ctx))
	})
	r.Route(http.MethodGet, apiv1.PathTerms, func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var requestModel model.Terms
		if err := mapstructure.Decode(request.QueryStringParameters, &requestModel); err != nil {
			return handleError(fmt.Errorf("error decoding request parameters: %v", err))
		}
		return marshalResponse(client.GetTerms(ctx, requestModel))
	})
	r.Route(http.MethodGet, apiv1.PathSubjects, func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var requestModel model.Subjects
		if err := mapstructure.Decode(request.QueryStringParameters, &requestModel); err != nil {
			return handleError(fmt.Errorf("error decoding request parameters: %v", err))
		}
		return marshalResponse(client.GetSubjects(ctx, requestModel))
	})
	r.Route(http.MethodGet, apiv1.PathCourses, func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var requestModel model.Courses
		if err := mapstructure.Decode(request.QueryStringParameters, &requestModel); err != nil {
			return handleError(fmt.Errorf("error decoding request parameters: %v", err))
		}
		return marshalResponse(client.GetCourses(ctx, requestModel))
	})
	r.Route(http.MethodGet, apiv1.PathSections, func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var requestModel model.Sections
		if err := mapstructure.Decode(request.QueryStringParameters, &requestModel); err != nil {
			return handleError(fmt.Errorf("error decoding request parameters: %v", err))
		}
		return marshalResponse(client.GetSections(ctx, requestModel))
	})
}

// Helper method to marshal data returned by the API client into an API Gateway Proxy Response
func marshalResponse(data interface{}, handlerErr error) (events.APIGatewayProxyResponse, error) {
	if handlerErr != nil {
		return handleError(handlerErr)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return handleError(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

// Helper method to return errors as an API Gateway Proxy Response
func handleError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       "the server has encountered an unexpected error",
	}, err
}

func initalizeLogger(ctx context.Context) (*logr.Logger, context.Context, error) {
	zLog, err := zap.NewProduction()
	if err != nil {
		return nil, nil, err
	}
	ctx = logr.NewContext(ctx, zapr.NewLogger(zLog))
	log, err := logr.FromContext(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &log, ctx, nil
}
