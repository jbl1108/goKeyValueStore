package input

type KeyValSyncInputPort interface {
	/*
		Retrieves the value associated with the provided key.
		Returns the value as a byte slice and an error if the operation fails.
	*/
	GetKey(key string) ([]byte, error)
	/*	Sets the value for the provided key.
		Takes a key as a string and a value as a byte slice.
		Returns an error if the operation fails.
	*/
	SetKey(key string, value []byte) error
}
