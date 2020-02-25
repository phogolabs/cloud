package pubsub_test

import (
	"context"
	"fmt"

	"github.com/phogolabs/cloud"
	"github.com/phogolabs/cloud/fake"
	"github.com/phogolabs/cloud/pubsub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Receiver", func() {
	var (
		receiver *pubsub.Receiver
		handler  *fake.EventHandler
	)

	BeforeEach(func() {
		handler = &fake.EventHandler{}

		receiver = &pubsub.Receiver{
			Handler: handler,
		}
	})

	It("receives the event successfully", func() {
		event := cloud.NewEvent()
		event.SetID("0001")
		event.SetType("dev.example.user.create")
		event.SetSource("http://example.com/services/my-service")

		msg := encode(&event)

		_, err := receiver.Receive(context.TODO(), msg)
		Expect(err).To(Succeed())

		Expect(handler.HandleCallCount()).To(Equal(1))

		ctx, args := handler.HandleArgsForCall(0)
		Expect(args).To(Equal(&event))

		tx := pubsub.TransportFromContext(ctx)
		Expect(tx).NotTo(BeNil())
		Expect(tx.Project).To(Equal("my-project"))
		Expect(tx.Subscription).To(Equal("my-subscription"))
	})

	Context("when the handler fails", func() {
		BeforeEach(func() {
			handler.HandleReturns(fmt.Errorf("oh no"))
		})

		It("returns an error", func() {
			event := cloud.NewEvent()
			event.SetID("0001")
			event.SetType("dev.example.user.create")
			event.SetSource("http://example.com/services/my-service")

			msg := encode(&event)

			_, err := receiver.Receive(context.TODO(), msg)
			Expect(err).To(MatchError("rpc error: code = Internal desc = oh no"))
		})
	})
})
