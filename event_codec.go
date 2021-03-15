package cloud

import (
	"context"
	"fmt"
	"reflect"

	"github.com/cloudevents/sdk-go/v2/event/datacodec"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func init() {
	datacodec.AddEncoder(ApplicationGRPCJSON, JSONPBEncode)
	datacodec.AddDecoder(ApplicationGRPCJSON, JSONPBDecode)
}

// JSONPBDecode takes `in` as []byte.
// If Event sent the payload as base64, Decoder assumes that `in` is the
// decoded base64 byte array.
func JSONPBDecode(ctx context.Context, in []byte, out interface{}) error {
	if in == nil {
		return nil
	}
	if out == nil {
		return fmt.Errorf("out is nil")
	}

	msg, ok := out.(proto.Message)
	if !ok {
		return fmt.Errorf("out is not proto.Message")
	}

	if err := protojson.Unmarshal(in, msg); err != nil {
		return fmt.Errorf("[jsonpb] found bytes \"%s\", but failed to unmarshal: %s", string(in), err.Error())
	}
	return nil
}

// Encode attempts to json.Marshal `in` into bytes. Encode will inspect `in`
// and returns `in` unmodified if it is detected that `in` is already a []byte;
// Or json.Marshal errors.
func JSONPBEncode(ctx context.Context, in interface{}) ([]byte, error) {
	if in == nil {
		return nil, nil
	}

	it := reflect.TypeOf(in)
	switch it.Kind() {
	case reflect.Slice:
		if it.Elem().Kind() == reflect.Uint8 {

			if b, ok := in.([]byte); ok && len(b) > 0 {
				// check to see if it is a pre-encoded byte string.
				if b[0] == byte('"') || b[0] == byte('{') || b[0] == byte('[') {
					return b, nil
				}
			}

		}
	}

	msg, ok := in.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("out is not proto.Message")
	}

	return protojson.Marshal(msg)
}
