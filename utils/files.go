package utils

import (
	"io"
	"log"
	"os"
)

func getWorkingDirectory() (string) {
	wd, err := os.Getwd()
		
	if (err != nil) {
		log.Fatal(err)
	}

	return wd
}

func DoesFileExist(filename string) (bool) {	
	file := getWorkingDirectory() + "/files/" + filename

	fileinfo, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !fileinfo.IsDir()
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
	file := getWorkingDirectory() + "/" + dir

	f, err := os.Create(file)
	
	if (err != nil) {
		log.Fatal(err)
	}

	io.Copy(f, data)
}

func CreateDir(dir string) {
	// store in ./files dir temporary
	err := os.Mkdir("files/" + dir , 0755)

	if err != nil {
		log.Fatal(err)
	}
}
