package fake

import (
	"github.com/phogolabs/cloud/provider/gcp"
)

// NewFakeErrorOption creates a new fake error option
func NewFakeErrorOption(err error) gcp.Option {
	fn := func(p *gcp.Protocol) error {
		return err
	}

	return fn
}
