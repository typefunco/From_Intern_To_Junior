# Interfaces

Итак, поговорим про интерфейсы. До этого момента я считал интерфейсы не супер полезными штуками. Ну вот что можно сделать? Определить интерфейс, а потом создать структуру, которая будет удовлетворять этому интерфейсу? Давайте тогда просто создадим структуру без каких-либо там интерфейсов. Вроде мысль здравая, но так скажем она не дальновидная и скорее всего вы не осознаете всей мощи интерфейсов в языке Go<br><br>

<h3>Представим, что мы хотим создать структуру и методы для неё</h3><br>
``` go
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
}
```

<br>Вроде выглядит неплохо. Создали человека, создали короля, все круто работает. Но все же тут есть пространство для улучшения. Давай
попробуем изменить код, чтобы он стал более правильным<br>

``` go
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
```
<br>Чем же лучше второй кусок кода? Тут все просто. У нас есть универсальный метод **ReadInfo**, который
позволяет просто передавать структуру, которая удовлетворяет интерфейсу **ActionsWithFile**. Всё. По сути
здесь мы сделали круто. Чем же круто? Ну если мы захотим поменять чтение файла с XML на JSON, то нам следует просто написать 
реализацию и поменять в функции **ReadInfo** На JsonFile и все. Теперь мы не привязаны к конкретной реализации, что
позволяет скрыть саму реализацию и дать абстракцию.