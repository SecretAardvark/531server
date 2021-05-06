package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/SecretAardvark/531webapp/lift"
)

func main() {
	http.HandleFunc("/", dataHandler)
	http.HandleFunc("/data", dataHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Form: ", r.Form) //show what data we received from client

	fmt.Println(r.PostFormValue("weight"))
	fmt.Println(r.PostFormValue("reps"))
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, t)
	} else {
		r.ParseForm()
		weight, err := strconv.Atoi(r.PostFormValue("weight"))
		if err != nil {
			fmt.Println(err)
		}
		reps, err := strconv.Atoi(r.PostFormValue("reps"))
		if err != nil {
			fmt.Println(err)
		}
		lift := &lift.Lift{Name: "OHP"}
		lift.GetOneRep(float32(weight), float32(reps))
		lift.GetTM()
		lift.GetCycle()

		fmt.Println(lift.OneRepMax)
		fmt.Println(lift.TrainingMax)
		fmt.Println(lift.Cycle)
		tmpl := template.Must(template.ParseFiles("data.html"))
		tmpl.Execute(w, lift)
	}

}

func cycleHandler(w http.ResponseWriter, r *http.Request) {
	//Get weight/reps variables from html
	//calculate training cycle as a go struct
	//write the struct to the html template
	//fmt.Println(r.Method)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, tmpl)
}

//
