/*
 * Copyright (c) 2020 by F4 Developer (Stanisław Kowański). This file is part of
 * dndEncounterCalculator project and is released under MIT License. For full license
 * details, search for LICENSE file in root project directory.
 */

/*
Package misc holds smaller, one-file parts of program
*/
package misc

import (
	"encoding/json"
	"github.com/kirsle/configdir"
	"io/ioutil"
	"os"
	"path/filepath"
)

//DataFile is a struct to manage files used to save data/values to disk
type DataFile struct {
	FileDir  string // file directory (by default app subdirectory in system application data directory)
	FileName string // requested file name WITH extension
	FullPath string // directory and filename combined
}

//NewDataFile is a constructor of DataFile struct.
//Accepts filename argument. Sets default file directory and combines a full path
func NewDataFile(fileName string) *DataFile {
	dataFile := DataFile{FileName: fileName}.loadDefaults()

	dataFile.FullPath = filepath.Join(dataFile.FileDir, dataFile.FileName)

	return dataFile
}

//loadDefaults gets default directory of files (app subdirectory in system application data directory)
func (receiver DataFile) loadDefaults() *DataFile {
	receiver.FileDir = configdir.LocalConfig("ddcalculator")
	err := configdir.MakePath(receiver.FileDir)
	Check(err)

	return &receiver
}

//LoadData reads data from .json file and tries to unmarshal it into the passed variable
func (receiver *DataFile) LoadData(target interface{}) {
	// Tries to open file from disk, defers closing the handle
	fh, err := os.Open(receiver.FullPath)
	Check(err)
	defer fh.Close()

	// Reads file
	byteValue, _ := ioutil.ReadAll(fh)

	// Tries to fit the file value into passed variable
	err = json.Unmarshal(byteValue, &target)
	if err != nil {
		target = nil
	}
}

//WriteData writes passed data to disk
func (receiver *DataFile) WriteData(d interface{}) {
	// Tries to open file from disk, defers closing the handle
	fh, err := os.Open(receiver.FullPath)
	Check(err)
	defer fh.Close()

	// Marshalds data into json
	byteValue, err := json.Marshal(d)
	Check(err)

	// Writes json data to file
	err = ioutil.WriteFile(receiver.FullPath, byteValue, 0644)
	Check(err)
}

//CheckFile checks if the file exists and creates an empty file if it doesn't
func (receiver *DataFile) CheckFile() *DataFile {
	// Check if file exists
	if _, err := os.Stat(receiver.FullPath); os.IsNotExist(err) {
		// Creates the file in the location, defers closing the handle
		fh, err := os.Create(receiver.FullPath)
		Check(err)
		defer fh.Close()
	}

	return receiver
}
