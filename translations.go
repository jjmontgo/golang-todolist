package main

import "golang-todolist/frame"

func init() {
	frame.Registry.Translations = map[string]map[string]string{
		"en": 	map[string]string{
			"addTodoLabel": "Add Todo",
			"backToTodolistsLabel": "Back to todo lists",
			"backToUsersLabel": "Back to users",
			"cancelLabel": "Cancel",
			"deleteConfirmation": "Are you sure?",
			"deleteLabel": "Delete",
			"editLabel": "Edit",
			"editTodoTitle": "Edit Todo",
			"editUserTitle": "Edit User",
			"emailAddressLabel": "Email Address",
			"emailLabel": "Email",
			"emailTodolistEmailLabel": "Who would you like to send the todolist to?",
			"emailTodolistTitle": "Email Todolist",
			"loginLabel": "Login",
			"logoutLabel": "Logout",
			"nameLabel": "Name",
			"newTodolistLabel": "New Todolist",
			"newUserLabel": "New User",
			"noImage": "No Image",
			"operationsLabel": "Operations",
			"passwordLabel": "Password",
			"saveLabel": "Save",
			"seeTodosLabel": "See Todos",
			"sendEmailLabel": "Send",
			"todoDescriptionLabel": "Type a description of the todo",
			"todolistEditTitle": "Edit Todolist",
			"todolistNameLabel": "Name of Todo List",
			"todolistsListLabel": "These are your todo lists:",
			"todolistsTitle": "Todo Lists",
			"usernameLabel": "Username",
			"usersTitle": "Users",
		},
		"fr": 	map[string]string{
			"addTodoLabel": "Add Todo",
			"backToTodolistsLabel": "Back to todo lists",
			"backToUsersLabel": "Back to users",
			"cancelLabel": "Cancel",
			"deleteConfirmation": "Are you sure?",
			"deleteLabel": "Delete",
			"editLabel": "Edit",
			"editTodoTitle": "Edit Todo",
			"editUserTitle": "Edit User",
			"emailAddressLabel": "Email Address",
			"emailLabel": "Email",
			"emailTodolistEmailLabel": "Who would you like to send the todolist to?",
			"emailTodolistTitle": "Email Todolist",
			"loginLabel": "Login",
			"logoutLabel": "Logout",
			"nameLabel": "Name",
			"newTodolistLabel": "New Todolist",
			"newUserLabel": "New User",
			"operationsLabel": "Operations",
			"noImage": "No Image",
			"passwordLabel": "Password",
			"saveLabel": "Save",
			"seeTodosLabel": "See Todos",
			"sendEmailLabel": "Send",
			"todoDescriptionLabel": "Type a description of the todo",
			"todolistEditTitle": "Edit Todolist",
			"todolistNameLabel": "Name of Todo List",
			"todolistsListLabel": "These are your todo lists:",
			"todolistsTitle": "Todo Lists",
			"usernameLabel": "Username",
			"usersTitle": "Users",
		},
	}
}

func t(term string, lang string) string {
	return frame.Registry.Translations[term][lang]
}
