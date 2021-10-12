package fake

import (
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/bxcodec/faker/v3"
	"github.com/googleapis/google-cloudevents-go/cloud/storage/v1"
)

// NewFakeStorageObjectData creates a new fake storage object data
func NewFakeStorageObjectData() *storage.StorageObjectData {
	return &storage.StorageObjectData{
		ID:              pointer.ToString(faker.UUIDHyphenated()),
		Name:            pointer.ToString(faker.CCNumber()),
		Bucket:          pointer.ToString(faker.CCNumber()),
		ContentLanguage: pointer.ToString("en-US"),
		ContentType:     pointer.ToString("application/octet-stream"),
		Size:            pointer.ToInt64(1024),
		Metadata:        make(map[string]string),
		MediaLink:       pointer.ToString(strings.ToLower(faker.URL())),
		SelfLink:        pointer.ToString(strings.ToLower(faker.URL())),
		Kind:            pointer.ToString("storage#object"),
	}
}
