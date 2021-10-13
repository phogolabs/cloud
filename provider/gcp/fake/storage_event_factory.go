package fake

import (
	"github.com/googleapis/google-cloudevents-go/cloud/storage/v1"
	"github.com/phogolabs/cloud/provider/gcp"

	. "github.com/onsi/gomega"
)

// ModifyStorageObjectDataFunc represents a function that modify the storage object
type ModifyStorageObjectDataFunc func(*storage.StorageObjectData)

// NewFakeStorageObjectFinalizeEvent creates an object finalized event.
func NewFakeStorageObjectFinalizedEvent(modifiers ...ModifyStorageObjectDataFunc) gcp.Event {
	data := NewFakeStorageObjectData()

	// modify the data
	for _, fn := range modifiers {
		fn(data)
	}

	event, err := gcp.NewStorageEvent("OBJECT_FINALIZE", data)
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectArchivedEvent creates an object archived event.
func NewFakeStorageObjectArchivedEvent(modifiers ...ModifyStorageObjectDataFunc) gcp.Event {
	data := NewFakeStorageObjectData()

	// modify the data
	for _, fn := range modifiers {
		fn(data)
	}

	event, err := gcp.NewStorageEvent("OBJECT_ARCHIVE", data)
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectDeletedEvent creates an object deleted event.
func NewFakeStorageObjectDeletedEvent(modifiers ...ModifyStorageObjectDataFunc) gcp.Event {
	data := NewFakeStorageObjectData()

	// modify the data
	for _, fn := range modifiers {
		fn(data)
	}

	event, err := gcp.NewStorageEvent("OBJECT_DELETE", data)
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectMetadataUpdatedEvent creates an object updated event.
func NewFakeStorageObjectMetadataUpdatedEvent(modifiers ...ModifyStorageObjectDataFunc) gcp.Event {
	data := NewFakeStorageObjectData()

	// modify the data
	for _, fn := range modifiers {
		fn(data)
	}

	event, err := gcp.NewStorageEvent("OBJECT_METADATA_UPDATE", data)
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}
