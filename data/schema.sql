DROP DATABASE IF EXISTS golang_todolist;
CREATE DATABASE golang_todolist;
CONNECT golang_todolist;
CREATE TABLE todo_list (
	id int unsigned auto_increment primary key,
	name varchar(500)
);
