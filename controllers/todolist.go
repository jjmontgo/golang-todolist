package controllers

import (
	// "fmt"
	// "log"
	// "os"
	"strings"
	"time"
	"golang-todolist/frame"
	"golang-todolist/frame/aws"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Todolist")

	this.IsAccessible = func(actionName string) bool {
		return this.UserIsLoggedIn()
	}

	this.Actions["Index"] = func() {
		var todoLists []model.TodoList
		this.DB().Preload("MediaAttachment").Find(&todoLists)
		this.RenderJSON("todolists", todoLists)
	}

	this.Actions["Edit"] = func() {
		id := this.ParamUint("id")
		var todoList model.TodoList
		// update an existing list
		if id != 0 {
			this.DB().First(&todoList, id)
		}
		// or create a new one
		if todoList.Id == 0 {
			todoList = model.TodoList{Name: "",}
		}
		this.RenderJSON("todolist", todoList)
	}

	this.Actions["Save"] = func() {
		list := model.TodoList{Id: this.ParamUint("id"), Name: this.Param("name")}
		list.Name = strings.Trim(this.Param("name"), " ")
		if list.Name == "" {
			this.RenderJsonError("message", "You must enter a name.")
			return
		}

		this.DB().Save(&list)
	}

	this.Actions["ImageForm"] = func() {
		// put object in root folder
		initialKeyPath := ""
		successActionStatus := "201"
		// redirect here after upload complete with URL params:
		// ?bucket=&key=&etag=
		// successActionRedirect := frame.AbsoluteURL("todolist_image_upload_complete", "id", this.Param("id"))
		successActionRedirect := "" // no redirect
		vars := aws.S3BrowserBasedUploadFormVariables(
			initialKeyPath,
			successActionStatus,
			successActionRedirect,
		)
		this.RenderJSON(
			"aws_upload_url", vars["aws_upload_url"],
			"key_path", vars["key_path"],
			"policy", vars["policy"],
			"success_action_status", successActionStatus,
			"success_action_redirect", successActionRedirect,
			"x_amz_algorithm", vars["x_amz_algorithm"],
			"x_amz_credential", vars["x_amz_credential"],
			"x_amz_date", vars["x_amz_date"],
			"x_amz_signature", vars["x_amz_signature"])
	}

	/**
	 * GET ?bucket=&key=&etag= from Amazon S3 Redirect
	 */
	this.Actions["ImageUploadComplete"] = func() {
		// is there already a media attachment?  delete it
		var existingMediaAttachment model.MediaAttachment
		this.DB().Where(&model.MediaAttachment{
			RefType: "todo_list",
			RefId: this.ParamUint("id"),
			Category: "main-image",
		}).First(&existingMediaAttachment)
		if existingMediaAttachment.Id != 0 {
			this.DB().Delete(&existingMediaAttachment)
		}

		this.DB().Create(&model.MediaAttachment{
			AwsS3ObjectKey: this.Param("key"),
			Category: "main-image",
			RefType: "todo_list",
			RefId: this.ParamUint("id"),
			CreatedAt: time.Now(), // YYYY-MM-DD HH:MM:SS
		})

		aws.ResizeS3ObjectImage(this.Param("key"), 150, 150)
	}

	this.Actions["Delete"] = func() {
		var todoList model.TodoList
		id := this.ParamUint("id")
		this.DB().First(&todoList, id)
		if todoList.Id != 0 {
			this.DB().Delete(&todoList)
		}
	}

	this.Actions["Email"] = func() {
		this.Render("todolist/email", "id", this.ParamUint("id"))
	}

	this.Actions["SendEmail"] = func() {
		body := "Here is your todolist:\n"
		var todoList model.TodoList
		id := this.ParamUint("id")
		this.DB().First(&todoList, id)

		var todos []model.Todo
		this.DB().Model(&todoList).Related(&todos)

		for _, todo := range todos {
			body += "* " + todo.Name + "\n"
		}
		this.Email(this.Param("email"), "Todolist", body, "")
	}
}
