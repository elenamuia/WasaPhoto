package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:id", rt.wrap(rt.setMyUserName))
	rt.router.DELETE("/users/:id", rt.wrap(rt.deleteMyProfile))
	rt.router.GET("/users/:id/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:id/mainstream", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:id/banned/:id", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:id/banned/:id", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:id/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:id/photos/:id", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:id/photos/:id", rt.wrap(rt.getPhoto))
	rt.router.PUT("/users/:id/photos/:id/comments/:id", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:id/photos/:id/comments/:id", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/users/:id/photos/:id/like/:id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:id/photos/:id/like/:id", rt.wrap(rt.unlikePhoto))
	rt.router.PUT("/users/:id/followed/:id", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:id/followed/:id", rt.wrap(rt.unfollowUser))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
