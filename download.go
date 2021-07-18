package fundconnext

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/codefin-stack/go-fundconnext/data"

	"github.com/mitchellh/mapstructure"
)

type Download struct {
	DataType    data.FundConnextFileType
	Header      DownloadHeader
	Body        []interface{}
	HeaderBytes []byte
	BodyBytes   []byte
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
		DataType: data.FundConnextFileTypeMapping[fileName],
	}
}

// Scan input will be the struct structure that you want output to be look like
func (d *Download) Scan(requestedStruct interface{}) error {
	return mapstructure.Decode(d.Body, &requestedStruct)
}

// Download FundConnext file (can save multiple paths)
func (f *FundConnext) Download(asOfDate string, fileType data.FundConnextFileType, optionalSavePath ...string) (result Download, err error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/files/%s/%s.zip", asOfDate, fileType.String())
	out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	if err != nil {
		return result, err
	}

	saveErr := saveFile(out, optionalSavePath)
	if saveErr != nil {
		log.Fatal(saveErr)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(out), int64(len(out)))
	if err != nil {
		log.Fatal(err)
	}

	if zipReader == nil || zipReader.File[0] == nil {
		return result, errors.New("empty zip file")
	}

	zipFile := zipReader.File[0]
	return parseDownloadFile(zipFile, fileType)
}

// Private Function only accessible for download func
func saveFile(data []byte, fullPath []string) error {

	if len(data) <= 0 || len(fullPath) <= 0 {
		log.Println(errors.New("(Skip) Either Empty Data or No Path"))
		return nil
	}

	for _, path := range fullPath {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			log.Println(errors.New("can't Save file"))
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			log.Println(errors.New("can't Save file"))
			return err
		}

		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, bytes.NewBuffer(data))
		if err != nil {
			log.Println(errors.New("can't Write file"))
			return err
		}
	}

	return nil
}

func parseDownloadFile(zipFile *zip.File, fileType data.FundConnextFileType) (Download, error) {
	f, err := zipFile.Open()
	if err != nil {
		return Download{}, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Get File Header
	scanner.Scan()
	headerString := scanner.Text()
	Header := strings.Split(headerString, "|")
	headerBytes := []byte(headerString)
	var Records int = StringToInt(Header[2])
	var DataStruct []interface{}
	var bodyBytes []byte

	for scanner.Scan() {
		lineString := scanner.Text()
		bodyBytes = append(bodyBytes, []byte(lineString)...)
		lineData := strings.Split(lineString, "|")
		downloadType := reflect.TypeOf(fileType.ModelType())
		reflectVal := reflect.New(downloadType)

		for key, value := range lineData {
			if reflectVal.Elem().Field(key).CanSet() {
				t, err := toReflectValue(value, reflectVal.Elem().Field(key).Interface())
				if err != nil {
					return Download{}, err
				}
				reflectVal.Elem().Field(key).Set(t)
			}
		}

		DataStruct = append(DataStruct, reflectVal.Elem().Interface())
	}

	var version string
	if len(Header) >= 4 {
		version = Header[3]
	}
	return Download{
		DataType: fileType,
		Header: DownloadHeader{
			Date:    Header[0],
			SA:      Header[1],
			Records: Records,
			Version: version,
		},
		Body:        DataStruct,
		HeaderBytes: headerBytes,
		BodyBytes:   bodyBytes,
	}, nil
}

func toReflectValue(text string, value interface{}) (reflect.Value, error) {
	var err error
	switch value.(type) {
	case *float32:
		var r float32
		if text != "" {
			s, err := strconv.ParseFloat(text, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = float32(s)
		}
		return reflect.ValueOf(&r), nil
	case float32:
		var r float32
		if text != "" {
			s, err := strconv.ParseFloat(text, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = float32(s)
		}
		return reflect.ValueOf(r), nil
	case *float64:
		var s float64
		if text != "" {
			s, err = strconv.ParseFloat(text, 64)

			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(&s), nil
	case float64:
		var s float64
		if text != "" {
			s, err = strconv.ParseFloat(text, 64)

			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(s), nil
	case *int8:
		var rd int8
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int8(s)
		}
		return reflect.ValueOf(&rd), nil
	case int8:
		var rd int8
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int8(s)
		}
		return reflect.ValueOf(rd), nil
	case *int16:
		var rd int16
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int16(s)
		}
		return reflect.ValueOf(&rd), nil
	case int16:
		var rd int16
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int16(s)
		}
		return reflect.ValueOf(rd), nil
	case *int32:
		var rd int32
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int32(s)
		}
		return reflect.ValueOf(&rd), nil
	case int32:
		var rd int32
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int32(s)
		}
		return reflect.ValueOf(rd), nil
	case *int64:
		var rd int64
		if text != "" {
			rd, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(&rd), nil
	case int64:
		var rd int64
		if text != "" {
			rd, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(rd), nil
	case *int:
		var rd int
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int(s)
		}
		return reflect.ValueOf(&rd), nil
	case int:
		var rd int
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint8:
		var rd uint8
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint8(s)
		}
		return reflect.ValueOf(&rd), nil
	case uint8:
		var rd uint8
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint8(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint16:
		var rd uint16
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint16(s)
		}
		return reflect.ValueOf(&rd), nil
	case uint16:
		var rd uint16
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint16(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint32:
		var rd uint32
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint32(s)
		}
		return reflect.ValueOf(&rd), nil
	case uint32:
		var rd uint32
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint32(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint64:
		var rd uint64
		if text != "" {
			rd, err = strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(&rd), nil
	case uint64:
		var rd uint64
		if text != "" {
			rd, err = strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(rd), nil
	case *uint:
		var rd uint
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint(s)
		}
		return reflect.ValueOf(&rd), nil
	case uint:
		var rd uint
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint(s)
		}
		return reflect.ValueOf(rd), nil
	case bool:
		var r bool
		if text != "" {
			switch text {
			case "Y", "y", "1", "T", "t":
				r = true
			case "N", "n", "0", "F", "f":
				r = false
			default:
				r = true
			}
		}
		return reflect.ValueOf(r), nil
	case *bool:
		var r bool
		if text != "" {
			switch text {
			case "Y", "y", "1", "T", "t":
				r = true
			case "N", "n", "0", "F", "f":
				r = false
			default:
				r = true
			}
			return reflect.ValueOf(&r), nil
		} else {
			return reflect.ValueOf(nil), nil
		}
	case *string:
		return reflect.ValueOf(&text), nil
	case string:
		return reflect.ValueOf(text), nil
	default:
		return reflect.ValueOf(nil), errors.New("no type conversion support")
	}
}
