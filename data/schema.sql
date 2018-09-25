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
CREATE TABLE user (
	id int unsigned auto_increment primary key,
	username varchar(100),
	email varchar(100),
	password_hash char(60),
	INDEX username(username),
	INDEX email(email),
	INDEX password_hash(password_hash)
);
CREATE TABLE media_attachment (
	id int unsigned auto_increment primary key,
	aws_s3_object_key char(150),
	category char(50), # 'main-image'
	ref_type char(50), # entity name 'todolist', 'todo'
	ref_id int unsigned,
	created_at datetime,
	UNIQUE aws_s3_object_key(aws_s3_object_key),
	INDEX ref(ref_type, ref_id),
	INDEX category(category)
);
