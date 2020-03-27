package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Server implements http.Handler. It validates incoming gitlab webhooks and
// then dispatches them to the appropriate plugins.
type Server struct{}

// ServeHTTP validates an incoming webhook and puts it into the event channel.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Error("Failed to get body for reporting the webhook event")
	}
	event, err := gitlab.ParseWebhook(gitlab.HookEventType(r), payload)
	if err != nil {
		logrus.WithError(err).Error("Failed to judge the webhook event")
	}

	fmt.Fprint(w, "Event received.")

	if err := s.demuxEvent(event, payload, r.Header); err != nil {
		logrus.WithError(err).Error("Error parsing event.")
	}

}

func (s *Server) demuxEvent(eventType interface{}, payload []byte, h http.Header) error {
	//todo we need to log what event be trigger
	//l := logrus.WithFields(
	//	logrus.Fields{
	//		"eventTypeField": eventType,
	//	},
	//)
	fmt.Println(reflect.TypeOf(eventType))
	switch eventType.(type) {
	case *gitlab.IssueCommentEvent:
		fmt.Println("Get IssueCommentEvent")

	case *gitlab.IssueEvent:
		fmt.Println("Get IssueEvent")

	case *gitlab.MergeEvent:
		fmt.Println("Get MergeEvent")

	}
	return nil
}
