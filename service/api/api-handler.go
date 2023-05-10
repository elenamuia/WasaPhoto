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
	rt.router.GET("/users/:userid/follower/", rt.wrap(rt.getFollower))
	rt.router.GET("/users/:userid/followed/", rt.wrap(rt.getFollowed))
	rt.router.PUT("/users/:userid/banned/:banneduser", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userid/banned/:banneduser", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:userid/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:userid/photos/:photoid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:userid/photos/:photoid", rt.wrap(rt.getPhoto))
	rt.router.POST("/users/:userid/photos/:photoid/comments/", rt.wrap(rt.commentPhoto))
	rt.router.GET("/users/:userid/photos/:photoid/listcomment/", rt.wrap(rt.getCommentList))
	rt.router.GET("/users/:userid/photos/:photoid/listlike/", rt.wrap(rt.getLikeList))
	rt.router.DELETE("/users/:userid/photos/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/users/:userid/photos/:photoid/like/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:userid/photos/:photoid/like/:likeid", rt.wrap(rt.unlikePhoto))
	rt.router.PUT("/users/:userid/followed/:followeduser", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userid/followed/:followeduser", rt.wrap(rt.unfollowUser))
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
