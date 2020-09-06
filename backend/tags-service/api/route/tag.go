package route

import (
	"github.com/gorilla/mux"
	"github.com/next-crisantojeronimo/aws-workshop/backend/tags-service/api/controller"
	constant "github.com/next-crisantojeronimo/aws-workshop/backend/tags-service/util"
)

// SetTagRoutes Set routes for tag resource
func SetTagRoutes(router *mux.Router) {
	subRouter := router.PathPrefix(constant.API_ROOT_PATH).Subrouter()

	subRouter.HandleFunc("/tags", controller.GetTags).Methods("GET")
	subRouter.HandleFunc("/tags", controller.PostTag).Methods("POST")
	subRouter.HandleFunc("/tags/{tagId}", controller.GetTag).Methods("GET")
	subRouter.HandleFunc("/tags/{tagId}", controller.DeleteTag).Methods("DELETE")
}
