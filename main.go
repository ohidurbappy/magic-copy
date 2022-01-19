package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func list_files(location string) []string {
	files, err := ioutil.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	var filenames []string

	for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	return filenames
}

func read_file(location string) string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func is_available(str string, list []string) bool {

	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func main() {

	go func() {

		text_files_location := "./texts"

		// is text_files_location available?
		if _, err := os.Stat(text_files_location); os.IsNotExist(err) {
			os.Mkdir(text_files_location, 0777)
		}

		text_files := list_files(text_files_location)

		for t := range text_files {
			text_files[t] = strings.Replace(text_files[t], ".txt", "", 1)
		}

		ch := clipboard.Watch(context.Background(), clipboard.FmtText)
		for data := range ch {

			str := ""
			if (string(data)) == "bappy" {
				str = `Ohidur Rahman Bappy`

			} else {

				if is_available(string(data), text_files) {
					str = read_file(text_files_location + "/" + string(data) + ".txt")
				}

			}

			if str != "" {

				clipboard.Write(clipboard.FmtText, []byte(str))
			}

		}

	}()

	loop := true
	for loop == true {

	}

	//Start listen hotkey event whenever you feel it is ready.
	// triggered := hk.Listen(context.Background())
	// for range triggered {
	// 	println("hotkey ctrl+s is triggered")
	// 	//clipboard.Write(clipboard.FmtText, []byte("text data"))
	// }

}
