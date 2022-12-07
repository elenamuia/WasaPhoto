package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login/", rt.wrap(rt.loginUser))
	rt.router.PUT("/user/:id/", rt.wrap(rt.updateUserName))
	rt.router.DELETE("/user/:id/", rt.wrap(rt.deleteProfile))
	rt.router.GET("/user/:id/", rt.wrap(rt.getProfile))
	rt.router.GET("/user/:id/mainstream", rt.wrap(rt.getMyMain))
	rt.router.PUT("/user/:id/banned/:id/", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:id/banned/:id/", rt.wrap(rt.unbanUser))
	rt.router.POST("/user/:id/photos", rt.wrap(rt.postPhoto))
	rt.router.DELETE("/user/:id/photos/:id/", rt.wrap(rt.deletePhoto))
	rt.router.GET("/user/:id/photos/:id/", rt.wrap(rt.getPhoto))
	rt.router.PUT("/user/:id/photos/:id/comment/:id/", rt.wrap(rt.addComment))
	rt.router.DELETE("/user/:id/photos/:id/comment/:id/", rt.wrap(rt.deleteComment))
	rt.router.PUT("/user/:id/photos/:id/like/:id/", rt.wrap(rt.addLike))
	rt.router.DELETE("/user/:id/photos/:id/comment/:id/", rt.wrap(rt.deleteLike))
	rt.router.PUT("/user/:id/photos/:id/followed/:id/", rt.wrap(rt.follow))
	rt.router.DELETE("/user/:id/photos/:id/followed/:id/", rt.wrap(rt.unfollow))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
