package usecases

import (
	"errors"
	"log"
	"strings"

	"github.com/jbl1108/goKeyValueStorage/usecases/datamodel"
	"github.com/jbl1108/goKeyValueStorage/usecases/ports/output"
)

type KeyValueAsyncHandling struct {
	outputPort output.KeyValueStorage
}

func NewKeyValueAsyncHandlingUseCase(outputPort output.KeyValueStorage) *KeyValueAsyncHandling {
	return &KeyValueAsyncHandling{outputPort: outputPort}
}

func (s *KeyValueAsyncHandling) HandleKeyValueMessage(message datamodel.Message) error {
	log.Printf("Handling key value message: %+v", message)
	if message.Data == nil {
		return errors.New("Message data is nil")
	}
	if _, ok := message.Data.(map[string]any); !ok {
		return errors.New("Message data is not a map[string]any")
	}
	prefix, err := s.getPrefix(message.Topic)
	if err != nil {
		return err
	}
	t := s.getKeyValue(message)
	key := prefix + ":" + t.Key
	return s.outputPort.Set(key, []byte(t.Value))
}

func (s *KeyValueAsyncHandling) getKeyValue(message datamodel.Message) *datamodel.KeyValue {
	m := message.Data.(map[string]any)
	t := new(datamodel.KeyValue)
	for k, v := range m {
		switch k {
		case "key":
			t.Key = k
			if str, ok := v.(string); ok {
				t.Value = str
			} else {
				log.Printf("Value for key 'key' is not a string: %v", v)
			}
		case "value":
			if str, ok := v.(string); ok {
				t.Value = str
			} else {
				log.Printf("Value for key 'value' is not a string: %v", v)
			}
		default:
			log.Printf("Unknown field in message data: %s", k)
		}
	}
	return t
}

func (*KeyValueAsyncHandling) getPrefix(topic string) (string, error) {
	parts := strings.Split(topic, "/")
	if len(parts) < 2 {
		return "", errors.New("Invalid topic format, expected 'keyvalue/{bucket}' got: " + topic)
	}
	if parts[0] != "keyvalue" {
		return "", errors.New("Invalid topic format, expected 'keyvalue/{bucket}' got: " + topic)
	}
	bucket := parts[1]
	return bucket, nil
}
