package connectors

import (
	"context"

	pb "github.com/open-cloud-initiative/fabric/gen/go/plugin/v1"
)

// Opts are the options.
type Opts struct {
	// Dry toggles the dry run mode.
	Dry bool
	// Force toggles the force mode.
	Force bool
	// Root runs the command as root.
	Root bool
}

// Connector is the interface for the connector plugins.
type Connector interface {
	// Migrate the plugin to the latest version
	Migrate(context.Context, *pb.Migrate_Request) error
	// Get the name of the plugin
	GetName(context.Context, *pb.GetName_Request) error
	// Get the current version of the plugin
	GetVersion(context.Context, *pb.GetVersion_Request) error
	// Send signal to flush and close open connections
	Close() error
}

var _ Connector = (*Unimplemented)(nil)

// Unimplemented is the default implementation.
type Unimplemented struct{}

// Migrate the plugin to the latest version
func (u *Unimplemented) Migrate(context.Context, *pb.Migrate_Request) error {
	return nil
}

// Get the name of the plugin
func (u *Unimplemented) GetName(context.Context, *pb.GetName_Request) error {
	return nil
}

// Get the current version of the plugin
func (u *Unimplemented) GetVersion(context.Context, *pb.GetVersion_Request) error {
	return nil
}

// Send signal to flush and close open connections
func (u *Unimplemented) Close() error {
	return nil
}
