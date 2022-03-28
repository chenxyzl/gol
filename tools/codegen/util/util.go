package util

import (
	"bufio"
	"bytes"
	"html/template"
	"os"
)

//模板{{}}填充值，用于tpl模板制作
func Render(path string, config map[string]interface{}) ([]byte, error) {
	var t, err = template.ParseFiles(path)
	if err != nil {

		return nil, err
	}

	bufer := bytes.NewBuffer(nil)
	defer bufer.Reset()
	err = t.Execute(bufer, config)
	if err != nil {

		return nil, err
	}

	return bufer.Bytes(), nil

}

func DirNames(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ReadLines(path string) ([]string, error) {

	lines := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {

		lines = append(lines, fscanner.Text())
	}
	return lines, nil
}
