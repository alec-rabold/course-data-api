package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	basePath    string
	routes      map[string]route
	routesMutex *sync.RWMutex
}

type route struct {
	methods      map[string]Handler
	methodsMutex *sync.RWMutex
}

type Handler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// NewRouter returns a new router instance
func NewRouter(basePath string) *Router {
	return &Router{
		basePath:    basePath,
		routes:      make(map[string]route),
		routesMutex: &sync.RWMutex{},
	}
}

// Route adds a new route to the router
func (r *Router) Route(method, path string, handler Handler) {
	r.routesMutex.Lock()
	defer r.routesMutex.Unlock()

	rt, exists := r.routes[path]
	if !exists {
		rt = route{
			methods:      make(map[string]Handler),
			methodsMutex: &sync.RWMutex{},
		}
	}
	rt.methodsMutex.Lock()
	defer rt.methodsMutex.Unlock()

	rt.methods[method] = handler
}

// Handle processes an incoming Lambda API Gateway Proxy Request
func (r *Router) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r.routesMutex.RLock()
	defer r.routesMutex.RUnlock()

	rt, exists := r.routes[strings.TrimPrefix(request.Path, r.basePath)]
	if !exists {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error: unsupported path (Path: '%s')", request.Path),
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	rt.methodsMutex.RLock()
	defer rt.methodsMutex.RUnlock()

	handler, exists := rt.methods[request.HTTPMethod]
	if !exists {
		if !exists {
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprintf("error: unsupported method for path (Path: '%s', Method: '%s')", request.Path, request.HTTPMethod),
				StatusCode: http.StatusBadRequest,
			}, nil
		}
	}

	return handler(ctx, request)
}
