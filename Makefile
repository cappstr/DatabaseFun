db_login:
	psql ${DATABSE_URL}

db_create_migration:
	migrate create -ext sql -dir migrations -seq $(name)

db_migration:
	migrate -database ${DATABSE_URL} -path migrations up
