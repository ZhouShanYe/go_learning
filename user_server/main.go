package main

import (
	"fmt"
	"html/template"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name   string
	Habits []string
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)

	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Println("value", strings.Join(value, ""))
	}
	fmt.Fprint(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("method", method)

	if method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		_ = r.ParseForm()
		fmt.Println("username:", r.Form["usename"])
		fmt.Println("password:", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "欢迎登陆， %s!", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "密码错误， 请重新登录")
		}
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-custom-Header", "custom")
	w.WriteHeader(201)

	user := &User{
		Name:   "zhoushanye",
		Habits: []string{"吃饭", "睡觉"},
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}
func main() {
	fmt.Println("123")

	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/write", write)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
