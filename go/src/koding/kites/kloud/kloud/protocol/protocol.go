package protocol

import "koding/kites/kloud/eventer"

// BuildOptions is passed to a build method. It contains all necessary
// informations.
type BuildOptions struct {
	// MachineId defines a unique ID in which the build informations are
	// fetched from. MachineId is used to gather the Username, ImageName,
	// InstanceName etc.. For example it could be a mongodb object id that
	// would point to a document that carries those informations or a key for a
	// key/value storage.
	MachineId string

	// Username defines the username on behalf the machine is being build.
	Username string

	// ImageName is used to build the machine based on this particular image.
	ImageName string

	// InstanceName is used to change the machine name (usually hostname). If
	// it's empty a unique name will be used.
	InstanceName string

	// Eventer pushes the latest events to the build event.
	Eventer eventer.Eventer
}

// BuildResponse should be returned from a Build method.
type BuildResponse struct {
	// InstanceName should define the name/hostname of the created machine. It
	// should be equal to the InstanceName that was passed via BuildOptions.
	InstanceName string

	// InstanceId should define a unique ID that defined the created machine.
	// It's different than the machineID and is usually an unique id which is
	// given by the third-party provider, for example DigitalOcean returns a
	// droplet Id.
	InstanceId int

	// KiteId should container the id that is deployed inside the machine.
	KiteId string

	// IpAddress defines the publid ip address of the running machine.
	IpAddress string
}

// Provider manages a machine. It is used to create and provision a single
// image or machine for a given Provider, to start/stop/destroy/restart a
// machine.
type Provider interface {
	// Prepare is responsible of configuring and initializing the builder and
	// validating the given configuration prior Build. Calling other methods
	// before Prepare should be forbidden.
	Prepare(...interface{}) error

	// Build is creating a image and a machine.
	Build(*BuildOptions) (*BuildResponse, error)

	// Start starts the machine
	Start(...interface{}) error

	// Stop stops the machine
	Stop(...interface{}) error

	// Restart restarts the machine
	Restart(...interface{}) error

	// Destroy destroys the machine
	Destroy(...interface{}) error

	// Info returns full information about a single machine
	Info(...interface{}) (interface{}, error)

	// Name returns the underlying provider type
	Name() string
}
