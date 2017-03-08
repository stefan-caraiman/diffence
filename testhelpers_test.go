package diffence

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func generateWantDiffFromFiles(filepathPath string) []DiffItem {
	want := []DiffItem{}
	filepathFile, _ := os.Open(filepathPath)
	r := io.MultiReader(
		filepathFile,
	)
	buffer := bytes.NewBuffer(make([]byte, 0))
	i := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		filepath := buffer.String()
		want = append(want, DiffItem{filepath: filepath})
		i++
		buffer.Reset()
	}
	return want
}

func getFixtureFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func getRuleFile(filename string) *[]Rule {
	rules, err := LoadRulesJSON(filename)
	if err != nil {
		panic(err)
	}
	return rules
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
