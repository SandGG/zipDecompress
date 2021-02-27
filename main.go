package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	readZip()
	unzip("./cont")
}

func readZip() {
	var zipR, err = zip.OpenReader("./files.zip")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range zipR.File {
		fmt.Println("File " + file.Name + " contains:")
		var r, err = file.Open()
		defer r.Close()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, r) //Show content
		fmt.Println()
	}
}

func unzip(dest string) {
	var zipR, err = zip.OpenReader("./files.zip")
	if err != nil {
		log.Fatal(err)
	}

	//Mkdir -> err if file exist, MkdirAll -> if file exist, does nothing
	os.MkdirAll(dest, 0755)

	for _, f := range zipR.File {
		extractFile(f, dest)
	}
	fmt.Println("-- Zip decompress successful --")
}

func extractFile(f *zip.File, dest string) {
	var rc, err = f.Open()
	defer rc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Joins any number of path elements into a single path
	var path = filepath.Join(dest, f.Name) //single path

	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode()) //Mode returns the permission
	} else {
		os.MkdirAll(filepath.Dir(path), f.Mode()) //Make directory in path
		var f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}

		//Copy of rc File in new file
		io.Copy(f, rc)
	}
}
