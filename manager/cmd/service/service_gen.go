// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/luw2007/thor/manager/pkg/endpoint"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Register":    {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Register", logger))},
		"Resource":    {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Resource", logger))},
		"ResourceAdd": {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ResourceAdd", logger))},
	}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Register", "Resource", "ResourceAdd"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
