INSERT INTO todo_list (name) VALUES
	('Things to do at home'),
	('Things to do at work')
;

INSERT INTO todo (todo_list_id, name) VALUES
	(1, 'Brush Teeth'),
	(1, 'Do the Laundry'),
	(1, 'Take out the Trash'),
	(2, 'Water the Vines'),
	(2, 'Listen to Matt'),
	(2, 'Nod at Will')
;

INSERT INTO user(username, email, password_hash) VALUES
	('jonathan', 'jjmontgo@gmail.com', '$2a$10$lschprsnzjE72bINba99h.HnQOSDbdiD8VpCQC8xiwKWO7eW0FlIO')
;
