package misc

import (
	"encoding/json"
	"github.com/kirsle/configdir"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DataFile struct {
	FileDir  string
	FileName string
	FullPath string
}

func NewDataFile(fileName string) *DataFile {
	dataFile := DataFile{FileName: fileName}.loadDefaults()

	dataFile.FullPath = filepath.Join(dataFile.FileDir, dataFile.FileName)

	return dataFile
}

func (receiver DataFile) loadDefaults() *DataFile {
	receiver.FileDir = configdir.LocalConfig("ddcalculator")
	err := configdir.MakePath(receiver.FileDir)
	Check(err)

	return &receiver
}

func (receiver *DataFile) LoadData(target interface{}) {
	fh, err := os.Open(receiver.FullPath)
	Check(err)
	defer fh.Close()

	byteValue, _ := ioutil.ReadAll(fh)
	err = json.Unmarshal(byteValue, &target)
	Check(err)
}

func (receiver *DataFile) WriteData(d interface{}) {
	fh, err := os.Open(receiver.FullPath)
	Check(err)
	defer fh.Close()

	byteValue, err := json.Marshal(d)
	Check(err)

	err = ioutil.WriteFile(receiver.FullPath, byteValue, 0644)
	Check(err)
}

func (receiver *DataFile) CheckFile() *DataFile {
	if _, err := os.Stat(receiver.FullPath); os.IsNotExist(err) {
		fh, err := os.Create(receiver.FullPath)
		Check(err)
		defer fh.Close()
	}

	return receiver
}
