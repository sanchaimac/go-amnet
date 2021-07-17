package models

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// Download ...
type Download struct {
	DataType FundConnextType
	Header   DownloadHeader
	Body     []interface{}
}

// DownloadHeader ...
type DownloadHeader struct {
	Date    string
	SA      string
	Records int
	Version string
}

// New ...
func (Download) New(fileName string) *Download {
	return &Download{
		DataType: FundConnextFileType[fileName],
	}
}

// Scan input will be the struct structure that you want output to be look like
func (d *Download) Scan(requestedStruct interface{}) (result []interface{}) {

	if reflect.ValueOf(requestedStruct).Kind() == reflect.Ptr {
		requestedStruct = reflect.ValueOf(requestedStruct).Elem().Interface()
	}
	for _, body := range d.Body {
		mapstructure.Decode(body, &requestedStruct)
		temp := requestedStruct

		result = append(result, temp)
	}
	return result
}
