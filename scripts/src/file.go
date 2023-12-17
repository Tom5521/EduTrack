package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func mkdir(dir string) {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		fmt.Println(err)
	}
}
func IsNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}
func IsExists(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsExist(err)
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

func DonwloadFile(path, url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error copying content to file:", err)
		return
	}
	fmt.Println("Download completed.")
}

func CopyFile(src, dest string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println(err)
	}
}
