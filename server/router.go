package server

import (
	"github.com/kbuzsaki/httpserv/http"
	"regexp"
)

type Route struct {
	Pattern string
	Handler Handler
}

func (route Route) Matches(request http.Request) bool {
	matched, err := regexp.MatchString(route.Pattern, request.Path)
	return matched && err == nil
}

type RoutingHandler struct {
	Routes  []Route
	Default Handler
}

func (rh *RoutingHandler) Handle(request http.Request) http.Response {
	for _, route := range rh.Routes {
		if route.Matches(request) {
			return route.Handler.Handle(request)
		}
	}

	return rh.Default.Handle(request)
}
