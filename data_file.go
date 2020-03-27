package main

import (
	"time"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LoadFile(file string) (lines []string) {
	lines = make([]string, 0, 1024)

	f, err := os.Open(file)
	if err != nil {
		LogError("open file " + file + ", err=" + err.Error())
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := readLine(reader)
		if err != nil {
			break
		}

		lines = append(lines, string(line))
	}

	LogInfo("loadFile " + file + " " + strconv.Itoa(len(lines)))
	return
}

func readLine(r *bufio.Reader) ([]byte, error) {
	line, isprefix, err := r.ReadLine()
	if !isprefix {
		return line, err
	} else {
		content := make([]byte, 0, len(line))
		content = append(content, line...)
		for isprefix && err == nil {
			line, isprefix, err = r.ReadLine()
			content = append(content, line...)
		}
		return content, err
	}
}

func LogError(str string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[error] ", str)
}

func LogInfo(str string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[info] ", str)
}
