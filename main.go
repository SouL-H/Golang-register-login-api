package main

import (
	"fmt"
	. "loginApi/WebServer"
	helper "loginApi/helpers"
	"net/http"
)

func errControl(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var hum Human

	uName, email, pass := "", "", ""
	mux := http.NewServeMux()
	mux.HandleFunc("/table", hum.ServeHTTP)
	
	mux.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		uName = r.FormValue("username")
		email = r.FormValue("email")
		pass = r.FormValue("pass")
		uNameCheck := helper.IsEmpty(uName)
		emaiCheck := helper.IsEmpty(email)
		passCheck := helper.IsEmpty(pass)
		if uNameCheck || emaiCheck || passCheck {
			fmt.Fprintf(w, "Please fill all the fields")
			return
		}
		fmt.Fprintf(w, "Welcome %s", uName)

	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pass = r.FormValue("pass")
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pass)
		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data")
			return
		}
		dbEmail := "abc@com"
		dbPass := "123"
		if email == dbEmail && pass == dbPass {
			fmt.Fprintf(w, "Login successful")
		} else {
			fmt.Fprintf(w, "Login failed")
		}
	})
	http.ListenAndServe(":8080", mux)
}
