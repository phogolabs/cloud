package fake

import (
	"github.com/phogolabs/cloud"

	. "github.com/onsi/gomega"
)

// NewFakeEventWith creates a new fake event with a given args
func NewFakeEventWith(args cloud.EventArgs) cloud.Event {
	eventArgs, err := cloud.NewEventWith(args)
	Expect(err).To(Succeed())

	return *eventArgs
}
