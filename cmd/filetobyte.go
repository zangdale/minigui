package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne"
)

//加载资源
//func LoadFile() {
//	//fyne.LoadResourceFromPath()
//	//fyne.LoadResourceFromURLString()
//	//fyne.NewStaticResource()
//}
func loadToByte(filename string) {
	fr, err := fyne.LoadResourceFromPath(filename)
	if err != nil {
		panic("Read Error in " + filename)
	}
	fmt.Println("Reading " + filename)
	fmt.Println(fr.Content())
}

func fileToData(filename string) string {
	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Read Error in " + filename)
	}
	fmt.Println("Reading " + filename + ": length was " + strconv.Itoa(len(fb)))

	buffer := bytes.NewBufferString("[]byte{")
	for i := range fb {
		buffer.WriteString(strconv.Itoa(int(fb[i])) + ",")
	}

	buffer.WriteString("}")

	fmt.Println(buffer.String())

	return buffer.String()
}
