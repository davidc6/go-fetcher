package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

const filesPath = "files"

func getWorkingDirectory() string {
	wd, err := os.Getwd()
		
	if (err != nil) {
		log.Fatal(err)
	}

	return wd
}

func CreateDirIfNotExists(relativePath string) bool {
	absolutePath := filepath.Join(getWorkingDirectory(), relativePath)
	fileinfo, err := os.Stat(absolutePath)

	if os.IsNotExist(err) {
		err := os.Mkdir(absolutePath, 0755)

		if err != nil {
			log.Fatal(err)
		}

		return true
	}

	return fileinfo.IsDir()
}


func DoesRegularFileExist(regularFile string) bool {
	absolutePath := filepath.Join(getWorkingDirectory(), filesPath, regularFile)
	fileinfo, err := os.Stat(absolutePath)

	if os.IsNotExist(err) {
		return false
	}

	return !fileinfo.IsDir()
}

func DoesDirExist(dir string) bool {
	absolutePath := filepath.Join(getWorkingDirectory(), filesPath, dir)
	fileinfo, err := os.Stat(absolutePath)

	if os.IsNotExist(err) {
		return false
	}

	return fileinfo.IsDir()
}

// Read the whole file into memory 
func StringToReader(filename string) (io.Reader) {
	wd := getWorkingDirectory() + "/files/" + filename

	f, err := os.Open(wd)
	
	if (err != nil) {
		log.Fatal(err)
	}
	
	return f
}

func SaveToDisk(dir string, data io.Reader) {
	file := filepath.Join(getWorkingDirectory(), filesPath, dir)

	f, err := os.Create(file)
	
	if (err != nil) {
		log.Fatal(err)
	}

	io.Copy(f, data)
}

func CreateDir(dir string) {
	absolutePath := filepath.Join(getWorkingDirectory(), filesPath, dir)
	err := os.Mkdir(absolutePath , 0755)

	if err != nil {
		log.Fatal(err)
	}
}
