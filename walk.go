package main

import (
	"log"
	"net/http"
	"html/template"

)

/*
extern void run();
extern void left();
extern void right();
extern void brake();
extern void back();

*/
import "C"


type Car struct {
	IP string
}

func (c *Car) Run() {
	C.run()
}
func (c*Car)Left() {
	C.left()
}
func (c*Car)Right() {
	C.right()
}
func (c*Car)Brake() {
	C.brake()
}
func (c*Car)Back() {
	C.back()
}

func main() {
	v1()
}
func v1() {
	//匿名函数注册路由
	http.HandleFunc("/", Deal)
	log.Println("starting service!")

	//log.Fatal输出后，会退出程序,执行os.Exit(1)
	log.Fatal(http.ListenAndServe(":80", nil))
}
func Deal(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("running"))
	var car Car
	car.IP = "192.168.12.1"
	tmpl, err := template.ParseFiles("page.html")
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.Execute(w, car); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	//gen cert
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	tForm := make(map[string]string)
	if r.Form["action"] == nil || len(r.Form["action"]) == 0 {
	}
	//return

	for a, b := range r.Form {
		if len(b) == 0 {
			//fmt.Println("a:",a,"b:","null")
			tForm[a] = ""
		} else {
			tForm[a] = b[0]
		}
	}
	switch tForm["action"] {
	case "forward":
		car.Run()
	case "back":
		car.Back()
	case "left":
		car.Left()
	case "right":
		car.Right()
	case "brake":
		car.Brake()
	}

}


