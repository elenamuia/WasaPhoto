package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login/", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:userid", rt.wrap(rt.setMyUserName))
	rt.router.DELETE("/users/:userid", rt.wrap(rt.deleteMyProfile))
	rt.router.GET("/users/:userid/profile/", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:userid/mainstream/", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userid/banned/:BannedId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userid/banned/:BannedId", rt.wrap(rt.unbanUser))
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/photos/:photoid", rt.wrap(rt.getPhoto))
	rt.router.POST("/photos/:photoid/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/photos/:photoid/like/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoid/like/:likeid", rt.wrap(rt.unlikePhoto))
	rt.router.PUT("/users/:userid/followed/:followedUserId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userid/followed/:followedUserId", rt.wrap(rt.unfollowUser))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
