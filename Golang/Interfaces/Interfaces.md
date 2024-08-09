# Interfaces

Итак, поговорим про интерфейсы. До этого момента я считал интерфейсы не супер полезными штуками. Ну вот что можно сделать? Определить интерфейс, а потом создать структуру, которая будет удовлетворять этому интерфейсу? Давайте тогда просто создадим структуру без каких-либо там интерфейсов. Вроде мысль здравая, но так скажем она не дальновидная и скорее всего вы не осознаете всей мощи интерфейсов в языке Go.

### Представим, что мы хотим создать структуру и методы для неё

```go
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
	fmt.Println("Saved XML file", filename)
}
func (xml *XMLFile) Read(filename string) {
	fmt.Println("Read XML file", filename)
}
func (xml *XMLFile) Update(filename string) {
	fmt.Println("Update XML file", filename)
}
func (xml *XMLFile) Delete(filename string) {
	fmt.Println("Delete XML file", filename)
}
func (xml *XMLFile) Show() {
	fmt.Println("File XML Path:", xml.path, "| File Weight:", xml.weight, "mb")
}

func main() {
	JsFile := JsonFile{
		path:   "usr/json",
		weight: 3.2,
	}
	XMLFile := XMLFile{
		path:   "usr/xml",
		weight: 5.8,
	}

	JsFile.Show()
	XMLFile.Show()
}
```
Вроде выглядит неплохо. Создали структуру для JSON и XML файлов, всё круто работает. Но всё же тут есть пространство для улучшения. Давайте попробуем изменить код, чтобы он стал более правильным.
```go
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
	fmt.Println("Saved XML file", filename)
}
func (xml *XMLFile) Read(filename string) {
	fmt.Println("Read XML file", filename)
}
func (xml *XMLFile) Update(filename string) {
	fmt.Println("Update XML file", filename)
}
func (xml *XMLFile) Delete(filename string) {
	fmt.Println("Delete XML file", filename)
}
func (xml *XMLFile) Show() {
	fmt.Println("File XML Path:", xml.path, "| File Weight:", xml.weight, "mb")
}

func ReadInfo(ac ActionsWithFile, filename string) {
	ac.Read(filename)
}

func main() {
	JsFile := JsonFile{
		path:   "usr/json",
		weight: 3.2,
	}
	XMLFile := XMLFile{
		path:   "usr/xml",
		weight: 5.8,
	}

	JsFile.Show()
	XMLFile.Show()

	ReadInfo(&JsFile, "usr/json")
	ReadInfo(&XMLFile, "usr/xml")
}
```