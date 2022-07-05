all: prepare
prepare:
	docker run -p 33066:3306 --name iman-test-db -e MYSQL_ROOT_PASSWORD=S3cret -d mysql:8
	echo "waiting for containers..."
	sleep 30
	mysql -uroot -P33066 -pS3cret -e "CREATE DATABASE authsvc; CREATE DATABASE postscrudsvc;"
