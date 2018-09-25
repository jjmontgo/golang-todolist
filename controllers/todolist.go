package controllers

import (
	// "fmt"
	// "log"
	"strings"
	"time"
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
		var todoLists []model.TodoList
		frame.GORM().Find(&todoLists)
		this.Render("todolist/index", "Results", todoLists)
	}

	this.Actions["Edit"] = func() {
		id := frame.StringToUint(this.Param("id"))
		var todoList model.TodoList
		// update an existing list
		if id != 0 {
			frame.GORM().First(&todoList, id)
		}
		// or create a new one
		if todoList.Id == 0 {
			todoList = model.TodoList{Name: "",}
		}
		this.Render("todolist/edit", "List", todoList)
	}

	this.Actions["Save"] = func() {
		list := model.TodoList{Id: frame.StringToUint(this.Param("id")), Name: this.Param("name")}
		list.Name = strings.Trim(this.Param("name"), " ")
		if list.Name == "" {
			this.Render("todolist/edit",
				"List", list,
				"Error", "You must enter a name.")
			return
		}

		frame.GORM().Save(&list)
		this.Redirect(frame.URL("index"))
	}

	this.Actions["ImageForm"] = func() {
		// put object in root folder
		initialKeyPath := ""
		successActionStatus := "201"
		// redirect here after upload complete with URL params:
		// ?bucket=&key=&etag=
		successActionRedirect := frame.AbsoluteURL("todolist_image_upload_complete", "id", this.Param("id"))
		vars := aws.S3BrowserBasedUploadFormVariables(
			initialKeyPath,
			successActionStatus,
			successActionRedirect,
		)
		this.Render("todolist/imageform",
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
		db := frame.GORM()
		var existingMediaAttachment model.MediaAttachment
		db.Where(&model.MediaAttachment{
			RefType: "todolist",
			RefId: frame.StringToUint(this.Param("id")),
			Category: "main-image",
		}).First(&existingMediaAttachment)
		if existingMediaAttachment.Id != 0 {
			db.Delete(existingMediaAttachment)
		}

		db.Create(&model.MediaAttachment{
			AwsS3ObjectKey: this.Param("key"),
			Category: "main-image",
			RefType: "todolist",
			RefId: frame.StringToUint(this.Param("id")),
			CreatedAt: time.Now(), // YYYY-MM-DD HH:MM:SS
		})

		this.Redirect(frame.URL("index"))
	}

	this.Actions["Delete"] = func() {
		db := frame.GORM()
		var todoList model.TodoList
		id := frame.StringToUint(this.Param("id"))
		db.First(&todoList, id)
		if todoList.Id != 0 {
			db.Delete(&todoList)
		}
		this.Redirect(frame.URL("index"))
	}

	this.Actions["Email"] = func() {
		this.Render("todolist/email", "id", frame.StringToUint(this.Param("id")))
	}

	this.Actions["SendEmail"] = func() {
		body := "Here is your todolist:\n"
		var todoList model.TodoList
		id := frame.StringToUint(this.Param("id"))
		db := frame.GORM()
		db.First(&todoList, id)

		var todos []model.Todo
		db.Model(&todoList).Related(&todos)

		for _, todo := range todos {
			body += "* " + todo.Name + "\n"
		}
		this.Email(this.Param("email"), "Todolist", body, "")
		this.Redirect(frame.URL("index"))
	}
}


