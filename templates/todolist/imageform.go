package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/imageform",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Upload Todolist Image</h1>

<a href="{{url "index"}}">Back to Todo Lists</a>

<form action="{{.aws_upload_url}}" method="POST" enctype="multipart/form-data">
	<input type="hidden" name="key" value="{{.key_path}}${filename}" />
	<input type="hidden" name="policy" value="{{.policy}}" />
	<input type="hidden" name="success_action_status" value="{{.success_action_status}}" />
	<input type="hidden" name="success_action_redirect" value="{{.success_action_redirect}}" />
	<input type="hidden" name="x-amz-algorithm" value="{{.x_amz_algorithm}}" />
	<input type="hidden" name="x-amz-credential" value="{{.x_amz_credential}}" />
	<input type="hidden" name="x-amz-date" value="{{.x_amz_date}}" />
	<input type="hidden" name="x-amz-signature" value="{{.x_amz_signature}}" />

	<div class="form-group">
		<input type="file" name="file">
	</div>
	<input type="submit" value="Save">
</form>
`})
}
