package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func DoesFileExist(filename string) (bool) {
	wd, e := os.Getwd()
		
	if (e != nil) {
		log.Fatal(e)
	}
	
	wd = wd + "/files/" + filename

	info, err := os.Stat(wd)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func File(filename string) (io.Reader) {
	wd, e := os.Getwd()
	
	if (e != nil) {
		log.Fatal(e)
	}

	wd = wd + "/files/" + filename
	
	data, e := ioutil.ReadFile(wd)
	s := string(data)
	
	return strings.NewReader(s)
}


func SaveToDisk(dir string, data io.Reader) {
	file, e := os.Getwd()
			
	if (e != nil) {
		log.Fatal(e)
	}
	
	file = file + "/" + dir

	f, err := os.Create(file)
	
	if (err != nil) {
		log.Fatal(err)
	}
	
	io.Copy(f, data)
}

func CreateDir(dir string) {
	err := os.Mkdir("files/" + dir , 0755)

	if err != nil {
		log.Fatal(err)
	}
}
