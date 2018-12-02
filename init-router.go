package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang-todolist/frame" // Dispatch(), Registry
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	frame.Registry.Router = r

	r.HandleFunc("/", frame.Dispatch("Todolist", "Index")).Name("index")

	r.HandleFunc("/todolist/new", frame.Dispatch("Todolist", "Edit")).Name("todolist_new")

	r.HandleFunc("/todolist/edit/{id:[0-9]+}", frame.Dispatch("Todolist", "Edit")).Name("todolist_edit")

	r.HandleFunc("/todolist/save", frame.Dispatch("Todolist", "Save")).
		Methods("POST").
		Name("todolist_save")

	r.HandleFunc("/todolist/image-form/{id:[0-9]+}", frame.Dispatch("Todolist", "ImageForm")).Name("todolist_image_form")

	r.HandleFunc("/todolist/image-upload-complete/{id:[0-9]+}", frame.Dispatch("Todolist", "ImageUploadComplete")).Name("todolist_image_upload_complete")

	r.HandleFunc("/todolist/delete/{id}", frame.Dispatch("Todolist", "Delete")).
		Methods("POST").
		Name("todolist_delete")

	r.HandleFunc("/todolist/email/{id}", frame.Dispatch("Todolist", "SendEmail")).
		Methods("POST").
		Name("todolist_send_email")

	r.HandleFunc("/todolist/{id:[0-9]+}", frame.Dispatch("Todo", "Index")).Name("todolist")

	r.HandleFunc("/todo/new/{todo_list_id:[0-9]+}", frame.Dispatch("Todo", "Edit")).Name("todo_new")

	r.HandleFunc("/todo/edit/{id:[0-9]+}", frame.Dispatch("Todo", "Edit")).Name("todo_edit")

	r.HandleFunc("/todo/save", frame.Dispatch("Todo", "Save")).
		Methods("POST").
		Name("todo_save")

	r.HandleFunc("/todo/delete/{id:[0-9]+}", frame.Dispatch("Todo", "Delete")).
		Methods("POST").
		Name("todo_delete")

	r.HandleFunc("/users", frame.Dispatch("User", "Index")).Name("users")

	r.HandleFunc("/user/new", frame.Dispatch("User", "Edit")).Name("user_new")

	r.HandleFunc("/user/edit/{id:[0-9]+}", frame.Dispatch("User", "Edit")).
		Name("user_edit")

	r.HandleFunc("/user/save", frame.Dispatch("User", "Save")).
		Methods("POST").
		Name("user_save")

	r.HandleFunc("/user/delete/{id:[0-9]+}", frame.Dispatch("User", "Delete")).
		Methods("POST").
		Name("user_delete")

	r.HandleFunc("/login", frame.Dispatch("Auth", "ValidateLogin")).
		Methods("POST").
		Name("login")

	r.HandleFunc("/logout", frame.Dispatch("Auth", "Logout")).Name("logout")

	r.HandleFunc("/urls", frame.Dispatch("App", "Urls")).Name("urls")

	r.HandleFunc("/translations", frame.Dispatch("App", "Translations")).
		Name("urls")

	return r
}
