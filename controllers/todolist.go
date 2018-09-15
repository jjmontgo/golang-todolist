package controllers

import (
	// "fmt"
	// "log"
	"encoding/base64"
	"os"
	"strings"
	"time"
	"golang-todolist/frame"
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
		awsUploadBucketName := os.Getenv("AWS_UPLOAD_BUCKET_NAME")
		awsUploadURL := "http://" + awsUploadBucketName + ".s3.amazonaws.com/"

		xAmzAlgorithm := "AWS4-HMAC-SHA256"

		successActionStatus := "204"
		successActionRedirect := frame.AbsoluteURL("index")

		currentTime := time.Now()
		date := currentTime.Format("2006") + currentTime.Format("01") + currentTime.Format("02")
		// xAmzDate := currentTime.UTC().Format(time.RFC3339)
		xAmzDate := date + "T000000Z"

		awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
		awsRegion := os.Getenv("AWS_REGION")
		xAmzCredential := awsAccessKeyId + "/" + date + "/" + awsRegion + "/s3/aws4_request"

		policy :=
			"{\"expiration\": \"2020-12-01T12:00:00.000Z\"," +
				"\"conditions\": [" +
					"{\"bucket\": \"" + awsUploadBucketName + "\" }," +
					"[\"starts-with\", \"$key\", \"\"]," +
					"{\"success-action-status\": \"" + successActionStatus + "\"}," +
					"{\"success-action-redirect\": \"" + successActionRedirect + "\"}," +
					"{\"x-amz-algorithm\": \"" + xAmzAlgorithm + "\"}," +
					"{\"x-amz-credential\": \"" + xAmzCredential + "\"}," +
					"{\"x-amz-date\": \"" + xAmzDate + "\"}," +
				"]" +
			"}"
		policy = base64.StdEncoding.EncodeToString([]byte(policy))

		xAmzSignature := ""
		this.Render("todolist/imageform",
			"aws_upload_url", awsUploadURL,
			"policy", policy,
			"success_action_status", successActionStatus,
			"success_action_redirect", successActionRedirect,
			"x_amz_algorithm", xAmzAlgorithm,
			"x_amz_credential", xAmzCredential,
			"x_amz_date", xAmzDate,
			"x_amz_signature", xAmzSignature)
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


