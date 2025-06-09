migrate create -seq -ext sql -dir ./migrations create_movies_table
migrate create -seq -ext sql -dir ./migrations add_movies_check_constraints
migrate create -seq -ext sql -dir ./migrations add_movies_indexes
migrate create -seq -ext sql -dir ./migrations create_users_table
migrate create -seq -ext sql -dir ./migrations create_tokens_table
migrate create -seq -ext sql -dir ./migrations add_permissions

migrate -path ./migrations -database='postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable' up


