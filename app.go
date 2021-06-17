package main

import (
	"moneygement-web/controller"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	/* var p model.Person
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&p); err != nil {
		// error
		log.Println("decode Error", err)
	}*/
	// fmt.Println(p)

	//msg := r.FormValue("msg")

	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//v := struct {
	//	Msg string `json:"msg"`
	//}{
	//	Msg: "hello",
	//}
	//if err := json.NewEncoder(w).Encode(v); err != nil {
	//	log.Println("Error: ", err)
	//}

	//tmpl := template.Must(template.New("sign").Parse("<html><body>{{.}}</body></html>"))
	//tmpl.Execute(w, msg)

}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/api/v1/user", middleWare(http.HandlerFunc(controller.UserHandler)))
	http.Handle("/api/v1/saving", middleWare(http.HandlerFunc(controller.SavingHandler)))
	http.ListenAndServe(":8080", nil)
}
