package fake

import (
	"github.com/phogolabs/cloud/provider/gcp"

	. "github.com/onsi/gomega"
)

// NewFakeStorageObjectFinalizeEvent creates an object finalized event.
func NewFakeStorageObjectFinalizedEvent() gcp.Event {
	event, err := gcp.NewStorageEvent("OBJECT_FINALIZE", NewFakeStorageObjectData())
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectArchivedEvent creates an object archived event.
func NewFakeStorageObjectArchivedEvent() gcp.Event {
	event, err := gcp.NewStorageEvent("OBJECT_ARCHIVE", NewFakeStorageObjectData())
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectDeletedEvent creates an object deleted event.
func NewFakeStorageObjectDeletedEvent() gcp.Event {
	event, err := gcp.NewStorageEvent("OBJECT_DELETE", NewFakeStorageObjectData())
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}

// NewFakeStorageObjectMetadataUpdatedEvent creates an object updated event.
func NewFakeStorageObjectMetadataUpdatedEvent() gcp.Event {
	event, err := gcp.NewStorageEvent("OBJECT_METADATA_UPDATE", NewFakeStorageObjectData())
	Expect(err).NotTo(HaveOccurred())
	Expect(event).NotTo(BeNil())
	return *event
}
