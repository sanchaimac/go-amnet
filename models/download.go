package models

import (
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
func (d *Download) Scan(requestedStruct interface{}) error {
	return mapstructure.Decode(d.Body, &requestedStruct)
}
