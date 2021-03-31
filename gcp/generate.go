package gcp

// we should generate the sdk
//go:generate clang-format -i proto/cloud_pubsub.proto
//go:generate buf generate
//go:generate buf lint

// we need to modify the swagger file
//go:generate strava apply
