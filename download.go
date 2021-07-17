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

	"github.com/codefin-stack/go-fundconnext/models"
)

// Download FundConnext file (can save multiple paths)
func (f *FundConnext) Download(asOfDate, fileName string, optionalSavePath ...string) (result models.Download, err error) {
	cfg := MakeAPICallerConfig(f)
	url := fmt.Sprintf("/api/files/%s/%s.zip", asOfDate, fileName)
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
	// readFile, err := readZipFile(zipFile)
	// if err != nil {
	// 	return result, err
	// }
	// parseDownloadFile(zipFile, fileName)
	// _ = unzippedFileBytes
	return parseDownloadFile(zipFile, fileName), nil
}

// Private Function only accessible for download func
func saveFile(data []byte, fullPath []string) error {

	if len(data) <= 0 || len(fullPath) <= 0 {
		log.Println(errors.New("(Skip) Either Empty Data or No Path"))
		return nil
	}

	for _, path := range fullPath {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			log.Println(errors.New("Can't Save file"))
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			log.Println(errors.New("Can't Save file"))
			return err
		}

		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, bytes.NewBuffer(data))
		if err != nil {
			log.Println(errors.New("Can't Write file"))
			return err
		}
	}

	return nil
}

func readZipFile(zf *zip.File) (io.ReadCloser, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return f, nil
}

func parseDownloadFile(zipFile *zip.File, fileName string) models.Download {
	f, err := zipFile.Open()
	if err != nil {
		return models.Download{}
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Get File Header
	scanner.Scan()
	Header := strings.Split(scanner.Text(), "|")

	var Records int = StringToInt(Header[2])
	var DataStruct []interface{}

	for scanner.Scan() {
		lineData := strings.Split(scanner.Text(), "|")
		downloadType := reflect.TypeOf(models.FundConnextFileType[fileName].ModelType())
		reflectVal := reflect.New(downloadType)

		// Original Code
		// for i := 0; i < reflectVal.Elem().NumField(); i++ {
		// 	if reflectVal.Elem().Field(i).CanSet() {
		// 		reflectVal.Elem().Field(i).Set(toReflectValue(lineData[i], reflectVal.Elem().Field(i).Interface()))
		// 	}
		// }

		for key, value := range lineData {
			if reflectVal.Elem().Field(key).CanSet() {
				reflectVal.Elem().Field(key).Set(toReflectValue(value, reflectVal.Elem().Field(key).Interface()))
			}
		}

		DataStruct = append(DataStruct, reflectVal.Elem().Interface().(models.FundNAV))
	}
	return models.Download{
		DataType: models.FundConnextFileType[fileName],
		Header: models.DownloadHeader{
			Date:    Header[0],
			SA:      Header[1],
			Records: Records,
			Version: Header[3],
		},
		Body: DataStruct,
	}
}

func toReflectValue(text string, value interface{}) reflect.Value {
	var err error
	switch value.(type) {
	case *float64:
		var s float64
		if text != "" {
			s, err = strconv.ParseFloat(text, 64)

			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&s)
	case float64:
		var s float64
		if text != "" {
			s, err = strconv.ParseFloat(text, 64)

			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&s)
	case *int64:
		var rd int64
		if text != "" {
			rd, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&rd)
	case int64:
		var rd int64
		if text != "" {
			rd, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&rd)
	case *uint64:
		var rd uint64
		if text != "" {
			rd, err = strconv.ParseUint(text, 10, 64)
			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&rd)
	case uint64:
		var rd uint64
		if text != "" {
			rd, err = strconv.ParseUint(text, 10, 64)
			if err != nil {
				panic("Fail to convert data")
			}
		}
		return reflect.ValueOf(&rd)
	case *string:
		return reflect.ValueOf(&text)
	case string:
		return reflect.ValueOf(&text)
	default:
		panic("match no type")
	}

}
