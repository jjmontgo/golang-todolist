DROP DATABASE IF EXISTS golang_todolist;
CREATE DATABASE golang_todolist;
CONNECT golang_todolist;
CREATE TABLE todo_list (
	id int unsigned auto_increment primary key,
	name varchar(500)
);
CREATE TABLE todo (
	id int unsigned auto_increment primary key,
	name varchar(500),
	todo_list_id int unsigned,
	INDEX todo_list_id(todo_list_id)
);
