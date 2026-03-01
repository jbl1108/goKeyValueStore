package input

import "github.com/jbl1108/goKeyValueStorage/usecases/datamodel"

type KeyValASyncInputPort interface {
	HandleKeyValueMessage(message datamodel.Message) error
}
