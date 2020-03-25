//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("Icon.png")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	buf.WriteString("package main\n")
	buf.WriteString("// generated by gen.go; DO NOT EDIT\n")
	buf.WriteString("var iconBytes = []byte{\n")
	for _, v := range data {
		fmt.Fprintf(&buf, "0x%02x,", v)
	}
	buf.WriteString("}")

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Error formatting generated code", err)
	}
	err = ioutil.WriteFile("generated.go", formatted, 0644)
	if err != nil {
		log.Fatal("Error writing blob file", err)
	}

}