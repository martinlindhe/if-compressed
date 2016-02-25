package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file  = kingpin.Arg("file", "File to compress.").Required().String()
	human = kingpin.Flag("human", "Human-readable output.").Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	if !exists(*file) {
		fmt.Println("ERROR file not found", *file)
		os.Exit(1)
	}

	presentResult("in   : ", getFileSize(*file))
	fmt.Println("")

	presentResult("gz  -5: ", compressGzip(*file, 5))
	presentResult("xz  -5: ", compressXz(*file, 5))
	presentResult("bro -5: ", compressBrotli(*file, 5))
	fmt.Println("")

	presentResult("gz  -9: ", compressGzip(*file, 9))
	presentResult("xz  -9: ", compressXz(*file, 9))
	presentResult("bro -9: ", compressBrotli(*file, 9))
}

func presentResult(title string, size int64) {
	if *human {
		fmt.Println(title, (size / 1024), "KiB")
	} else {
		fmt.Println(title, size)
	}
}

func compressBrotli(fileName string, compressionLevel int) int64 {

	outFileName := path.Join(os.TempDir(), randString(10)+".brot")

	cmd := exec.Command("bro", "--quality", fmt.Sprintf("%d", compressionLevel), "--input", fileName, "--output", outFileName)

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	if !exists(outFileName) {
		fmt.Println("ERROR FAIL")
		return 0
	}

	size := getFileSize(outFileName)

	os.Remove(outFileName)

	return size
}
func compressXz(fileName string, compressionLevel int) int64 {

	outFileName := path.Join(os.TempDir(), randString(10)+".xz")

	os.Setenv("XZ_OPT", fmt.Sprintf("-%d", compressionLevel))
	cmd := exec.Command("tar", "-cvJf", outFileName, fileName)

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	if !exists(outFileName) {
		fmt.Println("ERROR FAIL")
		return 0
	}

	size := getFileSize(outFileName)

	os.Remove(outFileName)

	return size
}

func compressGzip(fileName string, compressionLevel int) int64 {

	outFileName := path.Join(os.TempDir(), randString(10)+".gz")

	os.Setenv("GZIP", fmt.Sprintf("-%d", compressionLevel))
	cmd := exec.Command("tar", "-cvzf", outFileName, fileName)

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	if !exists(outFileName) {
		fmt.Println("ERROR FAIL")
		return 0
	}

	size := getFileSize(outFileName)

	os.Remove(outFileName)

	return size
}

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func getFileSize(filepath string) int64 {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	res := ""
	for i := 0; i < n; i++ {
		res += string(letterBytes[rand.Intn(len(letterBytes))])
	}
	return res
}
