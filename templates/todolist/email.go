package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/email",
		HasLayout: true,
		Template: `
<h1>Send Todolist</h1>
<p>Who would you like to send the todolist to?</p>
{{$id := uint_to_string .id}}
<form method="POST" action="{{url "todolist_email" "id" $id}}">
	<p>
		<input type="text" name="email" type="email" placeholder="Email Address" class="form-control" required>
	</p>
	<p>
		<button type="submit" class="btn btn-primary">Send</button>
		<a href="{{url "index"}}" class="btn btn-link">Cancel</a>
	</p>
</form>
`})
}
