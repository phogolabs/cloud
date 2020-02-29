package pubsub_test

import (
	"context"
	"log"
	"testing"

	pubsubevent "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	cloud "github.com/phogolabs/cloud"
	pubsub "github.com/phogolabs/cloud/pubsub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPubsub(t *testing.T) {
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "PubSub Suite")
}

func encode(event *cloud.Event) *pubsub.ReceivedMessage {
	codec := &pubsubevent.Codec{
		Encoding: pubsubevent.StructuredV1,
	}

	m, err := codec.Encode(context.TODO(), *event)
	Expect(err).To(BeNil())

	message, ok := m.(*pubsubevent.Message)
	Expect(ok).To(BeTrue())

	return &pubsub.ReceivedMessage{
		Message: &pubsub.Message{
			Attributes: message.Attributes,
			Data:       message.Data,
		},
		Subscription: "projects/my-project/subscriptions/my-subscription",
	}
}
