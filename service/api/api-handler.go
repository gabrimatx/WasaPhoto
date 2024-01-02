package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Login
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	// User
	rt.router.PUT("/users/:userId", rt.wrap(rt.setMyUserName))
	rt.router.DELETE("/users/:userId", rt.wrap(rt.deleteUser))
	rt.router.GET("/users/:userId/stream/", rt.wrap(rt.getMyStream))
	// Photos
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))
	// Comments
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	// Likes
	rt.router.PUT("/photos/:photoId/likes/:userId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId", rt.wrap(rt.unlikePhoto))
	// Follows
	rt.router.PUT("/users/:userId/follows/:followId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/follows/:followId", rt.wrap(rt.unfollowUser))
	// Bans
	rt.router.PUT("/users/:userId/bans/:banId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/bans/:banId", rt.wrap(rt.unbanUser))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
