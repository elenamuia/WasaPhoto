package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:userid", rt.wrap(rt.setMyUserName))
	rt.router.DELETE("/users/:userid", rt.wrap(rt.deleteMyProfile))
	rt.router.GET("/users/:userid/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:userid/mainstream", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userid/banned/:id", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userid/banned/:id", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:userid/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:userid/photos/:id", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:userid/photos/:id", rt.wrap(rt.getPhoto))
	rt.router.PUT("/users/:userid/photos/:id/comments/:id", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:userid/photos/:id/comments/:id", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/users/:userid/photos/:id/like/:id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:userid/photos/:id/like/:id", rt.wrap(rt.unlikePhoto))
	rt.router.PUT("/users/:userid/followed/:followedUserId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userid/followed/:followedUserId", rt.wrap(rt.unfollowUser))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
