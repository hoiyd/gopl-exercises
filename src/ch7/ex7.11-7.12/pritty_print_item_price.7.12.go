package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var listTemp = template.Must(template.New("").Parse(`
<table>
<tr style='text-align:left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{ range $key, $value := . }}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{ end }}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := listTemp.Execute(w, &db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to execute template: %q\n", err)
	}
}
