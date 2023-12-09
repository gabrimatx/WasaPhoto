package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	//Follows
	rt.router.PUT("/users/{userId}/follows/{followId}", rt.wrap(rt.followUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
