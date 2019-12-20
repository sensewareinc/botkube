package filters

import (
	"github.com/infracloudio/botkube/pkg/config"
	"github.com/infracloudio/botkube/pkg/events"
	"github.com/infracloudio/botkube/pkg/filterengine"
	log "github.com/infracloudio/botkube/pkg/logging"
)

// ImageTagChecker add recommendations to the event object if latest image tag is used in pod containers
type ErrorSuppressor struct {
	Description string
}

// Register filter
func init() {
	filterengine.DefaultFilterEngine.Register(ErrorSuppressor{
		Description: "Suppresses certain errors that are not really errors, such as HTTP 400 liveness probe failures",
	})
}

// Run filers and modifies event struct
func (f ErrorSuppressor) Run(object interface{}, event *events.Event) {
	if event.Kind != "Pod" || event.Type != config.ErrorEvent {
		return
	}

	for _, str := range event.Messages {
		if str == "Liveness probe failed: HTTP probe failed with statuscode: 400" {
			event.Skip = true
		} else {
			log.Logger.Info(str)
		}
	}
}

// Describe filter
func (f ErrorSuppressor) Describe() string {
	return f.Description
}
