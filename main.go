package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("GET")	
	route.HandleFunc("/project-detail/{id}", projectDetail).Methods("GET")	
	route.HandleFunc("/project", project).Methods("POST")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}


func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)

	data := map[string]interface{} {
		"Id": id,
	}


	tmpl.Execute(w, data)
}


func project(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project Name : " + r.PostForm.Get("input-project")) 
	fmt.Println("Start Date : " + r.PostForm.Get("input-start"))
	fmt.Println("End Date : " + r.PostForm.Get("input-end"))
	fmt.Println("Description : " + r.PostForm.Get("input-desc"))
	fmt.Println("Icon Node JS : " + r.PostForm.Get("node"))
	fmt.Println("Icon React JS : " + r.PostForm.Get("react"))
	fmt.Println("Icon Next JS : " + r.PostForm.Get("next"))
	fmt.Println("Icon TypeScript : " + r.PostForm.Get("type"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	
}