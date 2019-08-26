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

// DatabaseAlreadyInitialize is used when user try to initialize the database and
//it is already initialized
type DatabaseAlreadyInitialize struct{}

func (e *DatabaseAlreadyInitialize) Error() string {
	return "Error initializing service. Database is already initialized"
}

// DatabaseAlreadyInitializeError returns a new DatabaseAlreadyInitialize error
func DatabaseAlreadyInitializeError() error {
	return &DatabaseAlreadyInitialize{}
}

// IsDatabaseAlreadyInitializeError checks if the error is a DatabaseAlreadyInitialize error
func IsDatabaseAlreadyInitializeError(err error) bool {
	_, ok := err.(*DatabaseAlreadyInitialize)
	return ok
}

// NoDatabaseOptions is used when user try to get the database and it is
// not initialized
type NoDatabaseOptions struct{}

func (e *NoDatabaseOptions) Error() string {
	return "Cant't initilize Database. Database config is not supplied"
}

// NoDatabaseOptionsError returns a new NoDatabaseOptionsError error
func NoDatabaseOptionsError() error {
	return &NoDatabaseOptions{}
}

// IsNoDatabaseOptionsError checks if the error is a NoDatabaseOptions error
func IsNoDatabaseOptionsError(err error) bool {
	_, ok := err.(*NoDatabaseOptions)
	return ok
}

// FirebaseNotAlreadyInitilialized is used when user try to get the firebase app and it is
// not initialized
type FirebaseNotAlreadyInitilialized struct{}

func (e *FirebaseNotAlreadyInitilialized) Error() string {
	return "Cant't initilize Database. Database config is not supplied"
}

// FirebaseNotAlreadyInitilializedError returns a new FirebaseNotAlreadyInitilializedError error
func FirebaseNotAlreadyInitilializedError() error {
	return &FirebaseNotAlreadyInitilialized{}
}

// IsFirebaseNotAlreadyInitilializedError checks if the error is a FirebaseNotAlreadyInitilialized error
func IsFirebaseNotAlreadyInitilializedError(err error) bool {
	_, ok := err.(*FirebaseNotAlreadyInitilialized)
	return ok
}

// NoFirebaseOptions is used when user try to get the Firebase and it is
// not initialized
type NoFirebaseOptions struct{}

func (e *NoFirebaseOptions) Error() string {
	return "Cant't initilize Firebase. Firebase config is not supplied"
}

// NoFirebaseOptionsError returns a new NoFirebaseOptionsError error
func NoFirebaseOptionsError() error {
	return &NoFirebaseOptions{}
}

// IsNoFirebaseOptionsError checks if the error is a NoFirebaseOptions error
func IsNoFirebaseOptionsError(err error) bool {
	_, ok := err.(*NoFirebaseOptions)
	return ok
}
