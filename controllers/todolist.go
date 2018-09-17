package controllers

import (
	// "fmt"
	// "log"
	"strings"
	"golang-todolist/frame"
	"golang-todolist/frame/aws"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Todolist")

	this.IsAccessible = func(actionName string) bool {
		return frame.UserIsLoggedIn()
	}

	this.Actions["Index"] = func() {
		todoLists := model.FindTodolists()
		this.Render("todolist/index", "Results", todoLists)
	}

	this.Actions["Edit"] = func() {
		id := frame.StringToUint(this.Param("id"))
		var list *model.Todolist
		// update an existing list
		if id != 0 {
			list = model.FindTodolist("id", id)
		}
		// or create a new one
		if list == nil {
			list = &model.Todolist{Name: "",}
		}
		this.Render("todolist/edit", "List", list)
	}

	this.Actions["Save"] = func() {
		list := model.Todolist{Id: frame.StringToUint(this.Param("id")), Name: this.Param("name")}
		list.Name = strings.Trim(this.Param("name"), " ")
		if list.Name == "" {
			this.Render("todolist/edit",
				"List", list,
				"Error", "You must enter a name.")
			return
		}

		err := frame.SaveRecord(&list)
		if err != nil {
			this.Error(err)
			return
		}
		this.Redirect(frame.URL("index"))
	}

	this.Actions["ImageForm"] = func() {
		keyPath := "test/"
		successActionStatus := "201"
		successActionRedirect := frame.AbsoluteURL("index")
		vars := aws.S3BrowserBasedUploadFormVariables(keyPath, successActionStatus, successActionRedirect)
		this.Render("todolist/imageform",
			"aws_upload_url", vars["aws_upload_url"],
			"key_path", keyPath,
			"policy", vars["policy"],
			"success_action_status", successActionStatus,
			"success_action_redirect", successActionRedirect,
			"x_amz_algorithm", vars["x_amz_algorithm"],
			"x_amz_credential", vars["x_amz_credential"],
			"x_amz_date", vars["x_amz_date"],
			"x_amz_signature", vars["x_amz_signature"])
	}

	this.Actions["Delete"] = func() {
		todolist := model.FindTodolist("id", this.Param("id"))
		todolist.Delete()
		this.Redirect(frame.URL("index"))
	}

	this.Actions["Email"] = func() {
		this.Render("todolist/email", "id", frame.StringToUint(this.Param("id")))
	}

	this.Actions["SendEmail"] = func() {
		body := "Here is your todolist:\n"
		list := model.FindTodolist("id", this.Param("id"))
		todos := list.GetTodos()
		for _, todo := range todos {
			body += "* " + todo.Name + "\n"
		}
		this.Email(this.Param("email"), "Todolist", body, "")
		this.Redirect(frame.URL("index"))
	}
}


