source ./.env
mysql -p${MYSQL_PASSWORD} -u ${MYSQL_USERNAME} ${MYSQL_DB} < data/schema.sql
mysql -p${MYSQL_PASSWORD} -u ${MYSQL_USERNAME} ${MYSQL_DB} < data/test-data.sql

