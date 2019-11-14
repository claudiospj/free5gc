/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SMPolicy

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/npcf-smpolicycontrol/v1")
	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreateSmPolicy",
		strings.ToUpper("Post"),
		"/sm-policies",
		CreateSmPolicy,
	},

	{
		"DeleteSmPolicy",
		strings.ToUpper("Post"),
		"/sm-policies/:smPolicyId/delete",
		DeleteSmPolicy,
	},

	{
		"GetSmPolicy",
		strings.ToUpper("Get"),
		"/sm-policies/:smPolicyId",
		GetSmPolicy,
	},

	{
		"UpdateSmPolicy",
		strings.ToUpper("Post"),
		"/sm-policies/:smPolicyId/update",
		UpdateSmPolicy,
	},
}