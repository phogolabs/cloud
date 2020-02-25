package pubsub_test

import (
	"context"

	"cloud.google.com/go/pubsub/pstest"
	"gitlab.com/phogolabs/cloud"
	"gitlab.com/phogolabs/cloud/pubsub"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sender", func() {
	var (
		sender *pubsub.Sender
		topic  *pubsub.Topic
		client *pubsub.Client
		server *pstest.Server
	)

	BeforeEach(func() {
		var err error

		server = pstest.NewServer()

		conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
		Expect(err).To(BeNil())

		client, err = pubsub.NewClient(context.TODO(), "my-project", option.WithGRPCConn(conn))
		Expect(err).To(BeNil())

		sender = &pubsub.Sender{
			TopicID: "my-topic",
			Client:  client,
		}

		topic, err = client.CreateTopic(context.Background(), sender.TopicID)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		err := topic.Delete(context.Background())
		Expect(err).To(BeNil())

		Expect(server.Close()).To(Succeed())
	})

	It("sends an event successfully", func() {
		event := cloud.NewEvent()
		event.SetID("0001")
		event.SetType("dev.example.user.create")
		event.SetSource("http://example.com/services/my-service")

		Expect(sender.Send(context.TODO(), &event)).To(Succeed())
	})

	Context("when the event is not valid", func() {
		It("returns an error", func() {
			event := cloud.NewEvent()
			event.SetType("dev.example.user.create")
			event.SetSource("http://example.com/services/my-service")

			Expect(sender.Send(context.TODO(), &event)).To(MatchError("rpc error: code = InvalidArgument desc = id: MUST be a non-empty string"))
		})
	})

	Context("when sending the event fails", func() {
		BeforeEach(func() {
			sender.TopicID = "unknown"
		})

		It("returns an error", func() {
			event := cloud.NewEvent()
			event.SetID("0001")
			event.SetType("dev.example.user.create")
			event.SetSource("http://example.com/services/my-service")

			Expect(sender.Send(context.TODO(), &event)).To(MatchError(`rpc error: code = Internal desc = rpc error: code = NotFound desc = topic "projects/my-project/topics/unknown"`))
		})
	})
})