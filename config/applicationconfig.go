package config

import (
	"github.com/jbl1108/goKeyValueStorage/delivery"
	"github.com/jbl1108/goKeyValueStorage/repositories"
	"github.com/jbl1108/goKeyValueStorage/usecases"
	"github.com/jbl1108/goKeyValueStorage/usecases/ports/output"
)

type Application struct {
	outputPort             output.KeyValueStorage
	MQTTClient             *delivery.MQTTClient
	RestService            *delivery.KeyValueRestService
	storeTimeSeriesUseCase *usecases.KeyValueHandling
}

func NewApplication() Application {
	c := NewConfig()
	op := repositories.NewValkeyRepository(c.KeyValueUser(), c.KeyValuePassword(), c.KeyValueDBURL())
	su := usecases.NewKeyValueHandling(op)
	sd := delivery.NewKeyValueRestService(c.RestAddress(), su)
	mqttClient := delivery.NewMQTTClient(c.MQTTAddress(), c.MQTTUsername(), c.MQTTPassword(), "keyvalue/#", su)

	return Application{
		outputPort:             op,
		storeTimeSeriesUseCase: su,
		MQTTClient:             mqttClient,
		RestService:            sd,
	}
}
