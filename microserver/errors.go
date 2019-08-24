package microserver

// ServiceNotYetInitialize is used when user try to get the service and it is
// not initialized
type ServiceNotYetInitialize struct{}

func (e *ServiceNotYetInitialize) Error() string {
	return "Error getting service. Service is not yet initialized"
}

// ServiceNotYetInitializeError returns a new ServiceNotYetInitialize error
func ServiceNotYetInitializeError() error {
	return &ServiceNotYetInitialize{}
}

// ServiceAlreadyInitialize is used when user try to initialize the service and it is
// already initialized
type ServiceAlreadyInitialize struct{}

func (e *ServiceAlreadyInitialize) Error() string {
	return "Error initializing service. Service is already initialized"
}

// ServiceAlreadyInitializeError returns a new ServiceAlreadyInitialize error
func ServiceAlreadyInitializeError() error {
	return &ServiceAlreadyInitialize{}
}

// DatabaseNotYetInitialize is used when user try to get the database and it is
// not initialized
type DatabaseNotYetInitialize struct{}

func (e *DatabaseNotYetInitialize) Error() string {
	return "Error getting database. Database is not yet initialized"
}

// DatabaseNotYetInitializeError returns a new DatabaseNotYetInitialize error
func DatabaseNotYetInitializeError() error {
	return &DatabaseNotYetInitialize{}
}
