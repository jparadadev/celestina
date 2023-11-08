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

		var body interface{}

		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		subscriptions, subscriptionsExists := ctr.subscriptions[eventId]
		if subscriptionsExists {
			for subscriptor_name, subscriptor_url := range subscriptions {
				fmt.Println("Sending data to subscriptor:", subscriptor_name)
				res, err := http.Post(subscriptor_url, "application/json", ctx.Request.Body)
				if res != nil {
					fmt.Println("Subscriptor: ", subscriptor_name, ", Status code:", res.StatusCode)
				} else {
					fmt.Println("Subscriptor: ", subscriptor_name, " Failed with error: ", err)
				}
			}
		}

		ctx.String(http.StatusOK, "Ok!")
	}
}
