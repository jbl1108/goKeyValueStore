package usecases

import "github.com/jbl1108/goKeyValueStorage/usecases/ports/output"

type KeyValueSyncHandling struct {
	storage output.KeyValueStorage
}

func NewKeySyncHandling(storage output.KeyValueStorage) *KeyValueSyncHandling {
	return &KeyValueSyncHandling{storage: storage}
}

func (uc *KeyValueSyncHandling) SetKey(key string, data []byte) error {
	err := uc.storage.Open()
	if err != nil {
		return err
	}
	defer uc.storage.Close()

	return uc.storage.Set(key, data)
}
func (uc *KeyValueSyncHandling) GetKey(key string) ([]byte, error) {
	err := uc.storage.Open()
	if err != nil {
		return nil, err
	}
	defer uc.storage.Close()

	return uc.storage.Get(key)
}
