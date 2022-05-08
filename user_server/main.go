package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	Password string
	Habits   []string
}

var UserById = make(map[int]*User)
var UserByName = make(map[string][]*User)

var DB *sql.DB

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
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		user := User{1, r.Form.Get("username"), r.Form.Get("password"), []string{"学习"}}
		store(user)
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "欢迎登陆， %s!", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "密码错误， 请重新登录")
		}
	}
}

func store(user User) {
	UserById[user.Id] = &user
	UserByName[user.Name] = append(UserByName[user.Name], &user)
}

func loginInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for _, user := range UserByName[r.Form.Get("username")] {
		fmt.Fprintf(w, "%v", user)
	}
}
func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-custom-Header", "custom")
	w.WriteHeader(201)

	user := &User{
		Id:       0,
		Name:     "zhoushanye",
		Password: "123456",
		Habits:   []string{"吃饭", "睡觉"},
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}
func main() {
	fmt.Println("123")

	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login_info", loginInfo)
	http.HandleFunc("/write", write)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
