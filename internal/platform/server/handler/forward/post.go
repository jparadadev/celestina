package forward

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForwardController struct {
	subscriptions map[string][]string
}

func New(subscriptions map[string][]string) ForwardController {
	ctr := ForwardController{
		subscriptions: subscriptions,
	}
	return ctr
}

// CheckHandler returns an HTTP handler to perform forwarding connections.
func (ctr *ForwardController) PostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventId := ctx.Param("eventid")

		subscriptions, subscriptionsExists := ctr.subscriptions[eventId]
		if subscriptionsExists {
			for _, subscriptor := range subscriptions {
				fmt.Println("Sending to subscriptor:", subscriptor)
			}
		}

		var body map[interface{}]interface{}

		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.String(http.StatusOK, "Ok!")
	}
}
