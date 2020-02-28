package json_test

import (
	"context"
	"fmt"

	"github.com/phogolabs/cloud/fake"
	"github.com/phogolabs/cloud/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Receiver", func() {
	var (
		receiver *json.Receiver
		handler  *fake.EventHandler
	)

	BeforeEach(func() {
		handler = &fake.EventHandler{}

		receiver = &json.Receiver{
			Config: &json.ReceiverConfig{
				EventName:    "my-event",
				EventSubject: "UID",
				EventSource:  "example.com/service",
			},
			Handler: handler,
		}
	})

	It("receives the event successfully", func() {
		payload := json.Payload{
			"UID": "12345",
		}

		Expect(receiver.Receive(context.TODO(), payload)).To(Succeed())

		_, event := handler.HandleArgsForCall(0)
		Expect(event.ID()).NotTo(BeEmpty())
		Expect(event.Type()).To(Equal("my-event"))
		Expect(event.Source()).To(Equal("example.com/service"))
		Expect(event.Subject()).To(Equal("12345"))
	})

	Context("when the handler fails", func() {
		BeforeEach(func() {
			handler.HandleReturns(fmt.Errorf("oh no"))
		})

		It("returns an error", func() {
			payload := json.Payload{}
			Expect(receiver.Receive(context.TODO(), payload)).To(MatchError("oh no"))
		})
	})
})
