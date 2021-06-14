package main

import (
	"html/template"
	"moneygement-web/controller"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	/* var p model.Person
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&p); err != nil {
		// error
		log.Println("decode Error", err)
	}*/
	// fmt.Println(p)

	msg := r.FormValue("msg")

	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//v := struct {
	//	Msg string `json:"msg"`
	//}{
	//	Msg: "hello",
	//}
	//if err := json.NewEncoder(w).Encode(v); err != nil {
	//	log.Println("Error: ", err)
	//}

	tmpl := template.Must(template.New("sign").Parse("<html><body>{{.}}</body></html>"))
	tmpl.Execute(w, msg)

}

func main() {
	http.HandleFunc("/api/v1/user", controller.UserHandler)
	http.HandleFunc("/api/v1/saving", controller.SavingHandler)
	http.ListenAndServe(":8080", nil)
}
