CREATE TABLE session (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	session_data LONGBLOB,
	created_on TIMESTAMP default '0000-00-00 00:00:00',
	modified_on TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
	expires_on TIMESTAMP default '0000-00-00 00:00:00'
) ENGINE=InnoDB;
