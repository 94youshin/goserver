package store

var client Factory

// Factory defines the apiserver platform storage interface.
type Factory interface {
	Close() error
}

// Client return the sotre client instance.
func Client() Factory {
	return client
}

// SetClient set the apiserver store client.
func SetClient(factory Factory) {
	client = factory
}
