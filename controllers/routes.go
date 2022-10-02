package controllers

import "github.com/gin-gonic/gin"

//
// DoRoutes - Do Routes
//
func (t *Controller) DoRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")

	apiV1.POST("/customers/invite", t.InviteCustomers)
}
