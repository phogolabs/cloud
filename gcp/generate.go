package gcp

import "github.com/cloudevents/sdk-go/v2/binding/format"

func init() {
	format.Add(&PubsubFormat{})
	format.Add(&StorageFormat{})
}

// we should generate the sdk
//go:generate clang-format -i proto/cloud_pubsub.proto
//go:generate clang-format -i proto/cloud_storage.proto
//go:generate buf generate
//go:generate buf lint

// we need to modify the swagger file
//go:generate strava apply
