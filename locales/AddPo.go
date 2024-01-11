//go:build AddPo
// +build AddPo

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"gopkg.in/yaml.v3"
)

func GetFilesInDirectory(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func AddLineToFile(filename, line string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintln(writer, line)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(f string) map[string]string {
	var ret map[string]string
	file, _ := os.ReadFile(f)
	yaml.Unmarshal(file, &ret)
	return ret
}

const Template string = `

#: %s
msgid "%s"
msgstr ""
`

func main() {
	data := ReadFile("locales/last-add.yml")
	dirs, err := GetFilesInDirectory("locales/po/")
	if err != nil {
		messages.FatalError(err)
	}
	for _, file := range dirs {
		err := AddLineToFile(file, fmt.Sprintf(Template, data["route"], data["msgid"]))
		if err != nil {
			messages.FatalError(err)
		}
	}
}
