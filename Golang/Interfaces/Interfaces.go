package main

import "fmt"

type ActionsWithFile interface {
	Save(filename string)
	Read(filename string)
	Update(filename string)
	Delete(filename string)
	Show()
}

type JsonFile struct {
	path   string
	weight float32
}

func (js *JsonFile) Save(filename string) {
	fmt.Println("Saved JSON file", filename)
}
func (js *JsonFile) Read(filename string) {
	fmt.Println("Read JSON file", filename)
}
func (js *JsonFile) Update(filename string) {
	fmt.Println("Update JSON file", filename)
}
func (js *JsonFile) Delete(filename string) {
	fmt.Println("Delete JSON file", filename)
}
func (js *JsonFile) Show() {
	fmt.Println("File JSON Path:", js.path, "| File Weight:", js.weight, "mb")
}

type XMLFile struct {
	path   string
	weight float32
}

func (xml *XMLFile) Save(filename string) {
	fmt.Println("Saved XMl file", filename)
}
func (xml *XMLFile) Read(filename string) {
	fmt.Println("Read XMl file", filename)
}
func (xml *XMLFile) Update(filename string) {
	fmt.Println("Update XMl file", filename)
}
func (xml *XMLFile) Delete(filename string) {
	fmt.Println("Delete XMl file", filename)
}
func (xml *XMLFile) Show() {
	fmt.Println("File XMl Path:", xml.path, "| File Weight:", xml.weight, "mb")
}

func ReadInfo(ac ActionsWithFile, filename string) {
	ac.Read(filename)
}

func main() {
	JsFile := JsonFile{
		path:   "usr/json",
		weight: 3.2,
	}
	XMlFile := XMLFile{
		path:   "usr/xml",
		weight: 5.8,
	}

	JsFile.Show()
	XMlFile.Show()

	ReadInfo(&JsFile, "usr/json")
	ReadInfo(&XMlFile, "usr/xml")
}
