package config

import (
	"github.com/jbl1108/goKeyValueStorage/delivery"
	"github.com/jbl1108/goKeyValueStorage/repositories"
	"github.com/jbl1108/goKeyValueStorage/usecases"
	"github.com/jbl1108/goKeyValueStorage/usecases/ports/output"
)

type Application struct {
	outputPort                  output.KeyValueStorage
	MQTTClient                  *delivery.MQTTClient
	RestService                 *delivery.KeyValueRestService
	keyValueAsyncStorageUsecase *usecases.KeyValueAsyncHandling
	storeTimeSeriesUseCase      *usecases.KeyValueSyncHandling
}

func NewApplication() Application {
	c := NewConfig()
	op := repositories.NewValkeyRepository(c.KeyValueUser(), c.KeyValuePassword(), c.KeyValueDBURL())
	su := usecases.NewKeySyncHandling(op)
	sd := delivery.NewKeyValueRestService(c.RestAddress(), su)
	au := usecases.NewKeyValueAsyncHandlingUseCase(op)
	mqttClient := delivery.NewMQTTClient(c.MQTTAddress(), c.MQTTUsername(), c.MQTTPassword(), "keyvalue/#", au)

	return Application{
		keyValueAsyncStorageUsecase: au,
		outputPort:                  op,
		storeTimeSeriesUseCase:      su,
		MQTTClient:                  mqttClient,
		RestService:                 sd,
	}
}
