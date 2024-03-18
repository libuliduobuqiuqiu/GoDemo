package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"sunrun/internal/gotour/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temprature")

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type fileHandler interface {
	open() int
	close() int
}

type fileContext struct {
	Width  int
	Height int
}

func (f *fileContext) open() int {
	return f.Width * f.Height
}

func (f *fileContext) close() int {
	return f.Width / f.Height
}

type tempFile struct {
	length int
	fileContext
}

func main() {
	names := []string{"linshukai", "zhangsan", "wangwu", "meisi", "cluo"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
	f := tempFile{length: 11, fileContext: fileContext{Width: 12, Height: 13}}
	fmt.Println(f.open())
	fmt.Println(f.close())

	var w io.Writer
	fmt.Printf("%T\n", w)

	w = os.Stdout
	w.Write([]byte("hello\n"))
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)
	w.Write([]byte("hello\n"))
	fmt.Printf("%T\n", w)
}
