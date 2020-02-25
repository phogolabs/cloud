package cloud_test

import (
	"context"
	"fmt"

	"github.com/phogolabs/cloud"
	"github.com/phogolabs/cloud/fake"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompositeEventTypeHandler", func() {
	var (
		composite *cloud.CompositeEventTypeHandler
		handler   *fake.EventHandler
	)

	BeforeEach(func() {
		handler = &fake.EventHandler{}

		composite = &cloud.CompositeEventTypeHandler{
			"dev.cliche.contact.create": handler,
		}
	})

	It("handles event successfully", func() {
		event := cloud.NewEvent()
		event.SetType("dev.cliche.contact.create")

		Expect(composite.Handle(context.TODO(), &event)).To(Succeed())
		Expect(handler.HandleCallCount()).To(Equal(1))

		_, args := handler.HandleArgsForCall(0)
		Expect(args).To(Equal(&event))
	})

	Context("when the handler fails", func() {
		BeforeEach(func() {
			handler.HandleReturns(fmt.Errorf("oh no"))
		})

		It("returns the error", func() {
			event := cloud.NewEvent()
			event.SetType("dev.cliche.contact.create")

			Expect(composite.Handle(context.TODO(), &event)).To(MatchError("oh no"))
		})
	})

	Context("when the event type is unknown", func() {
		It("does not handle the request", func() {
			event := cloud.NewEvent()
			event.SetType("dev.cliche.contact.update")

			Expect(composite.Handle(context.TODO(), &event)).To(Succeed())
		})
	})
})

var _ = Describe("CompositeEventSourceHandler", func() {
	var (
		composite *cloud.CompositeEventSourceHandler
		handler   *fake.EventHandler
	)

	BeforeEach(func() {
		handler = &fake.EventHandler{}

		composite = &cloud.CompositeEventSourceHandler{
			"http://example.com/services/my-service": handler,
		}
	})

	It("handles event successfully", func() {
		event := cloud.NewEvent()
		event.SetSource("http://example.com/services/my-service")

		Expect(composite.Handle(context.TODO(), &event)).To(Succeed())
		Expect(handler.HandleCallCount()).To(Equal(1))

		_, args := handler.HandleArgsForCall(0)
		Expect(args).To(Equal(&event))
	})

	Context("when the handler fails", func() {
		BeforeEach(func() {
			handler.HandleReturns(fmt.Errorf("oh no"))
		})

		It("returns the error", func() {
			event := cloud.NewEvent()
			event.SetSource("http://example.com/services/my-service")

			Expect(composite.Handle(context.TODO(), &event)).To(MatchError("oh no"))
		})
	})

	Context("when the event type is unknown", func() {
		It("does not handle the request", func() {
			event := cloud.NewEvent()
			event.SetSource("http://example.com/services/other-service")

			Expect(composite.Handle(context.TODO(), &event)).To(Succeed())
		})
	})
})

var _ = Describe("CompositeEventHandler", func() {
	var (
		multiplex *cloud.CompositeEventHandler
		handler   *fake.EventHandler
	)

	BeforeEach(func() {
		handler = &fake.EventHandler{}

		multiplex = &cloud.CompositeEventHandler{
			handler,
		}
	})

	It("handles event successfully", func() {
		event := cloud.NewEvent()
		event.SetType("dev.cliche.contact.create")

		Expect(multiplex.Handle(context.TODO(), &event)).To(Succeed())
		Expect(handler.HandleCallCount()).To(Equal(1))

		_, args := handler.HandleArgsForCall(0)
		Expect(args).To(Equal(&event))
	})

	Context("when the handler fails", func() {
		BeforeEach(func() {
			handler.HandleReturns(fmt.Errorf("oh no"))
		})

		It("returns the error", func() {
			event := cloud.NewEvent()
			event.SetType("dev.cliche.contact.create")

			Expect(multiplex.Handle(context.TODO(), &event)).To(MatchError("oh no"))
		})
	})
})
