package main

import "flag"
import "log"
import "io"
import "io/ioutil"
import "encoding/json"
import "os"
import "text/template"

var FLAG_TPL_PATH string
var FLAG_JSON_PATH string

func init() {
	flag.StringVar(
		&FLAG_TPL_PATH,
		"tpl",
		"",
		"Specified the path of go-style template.",
	)
	flag.StringVar(
		&FLAG_JSON_PATH,
		"json",
		"stdin",
		"Specified the path of json input.",
	)
	flag.Parse()
	if FLAG_TPL_PATH == "" {
		log.Fatal("FLAG_TPL_PATH cannot be empty.")
	}
}

func getTpl(tpl_file_path string) *template.Template {
	tpl, err := ioutil.ReadFile(tpl_file_path)
	if err != nil {
		log.Fatal("getTpl :: open tpl", err)
	}
	ret, err := template.New("").Parse(string(tpl))
	if err != nil {
		log.Fatal("getTpl :: parse tpl", err)
	}
	return ret
}

func getJson(json_file_path string) string {
	var json_bytes []byte
	var err error
	if json_file_path == "stdin" {
		json_bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		json_bytes, err = ioutil.ReadFile(json_file_path)
	}
	if err != nil {
		log.Fatal("getJson :: ", err)
	}
	return string(json_bytes)
}

func renderHtml(tpl *template.Template, json_str string, wr io.Writer) {
	obj := new(interface{})
	err := json.Unmarshal([]byte(json_str), obj)
	if err != nil {
		log.Fatal("renderHtml :: ", err)
	}
	err = tpl.Execute(wr, obj)
	if err != nil {
		log.Fatal("renderHtml :: ", err)
	}
}

func main() {
	renderHtml(getTpl(FLAG_TPL_PATH), getJson(FLAG_JSON_PATH), os.Stdout)
}
