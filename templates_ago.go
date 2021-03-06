package templates_ago

import (
    "strings"
    "html/template"
    "runtime"
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
    "encoding/json"
    "bytes"
)

type Templates map[string]*template.Template

func (self Templates) Execute(w http.ResponseWriter, data interface{}) error {
    //Use interface{} case type check?
    //make check if template exists
    return self[FuncName(2) + ".html"].ExecuteTemplate(w, "base", data)
}

func FuncName(level int) string {
    pc, _, _, ok := runtime.Caller(level)
    if !ok { return "Unknown" }
    fnName := runtime.FuncForPC(pc)
    if fnName == nil { return "Anon" }
    fnNameParts := strings.Split(fnName.Name(), ".")
    return strings.ToLower(fnNameParts[len(fnNameParts)-1])
}

func TplName() string {
   return fmt.Sprintf("%s%s", FuncName(2), ".html")
}

func LoadTemplates(tplPath string, tpls Templates) {
    if _, err := os.Stat("templates/base.html"); os.IsNotExist(err) {
        panic("Missing templates/base.html file")
    }
    files, err := ioutil.ReadDir(tplPath)
    if err != nil { panic(err) }
    for _, file := range files {
        if file.Name() == "base.html" { continue }
        // some error handling would be nice here
        tpls[strings.ToLower(file.Name())] = template.Must(template.New(strings.ToLower(file.Name())).Funcs(template.FuncMap{"toJson": ToJson}).ParseFiles(tplPath + file.Name(), tplPath + "base.html"))
    }
}

func NewTemplates() Templates {
    return make(map[string]*template.Template)
}

func ToJson(input interface{}) string {
  b := new(bytes.Buffer)
  enc := json.NewEncoder(b)
  enc.Encode(input)
  return b.String()
}


