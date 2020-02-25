package pubsub

// download depencies
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto ./google/api
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/resource.proto ./google/api
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto ./google/api
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto ./google/api
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/client.proto ./google/api
//go:generate go-getter https://raw.githubusercontent.com/googleapis/googleapis/master/google/pubsub/v1/pubsub.proto ./google/pubsub/v1

// genarate boilerplate code
//go:generate protoc -I . --go_out=plugins=grpc:$GOPATH/src/. spec.proto
