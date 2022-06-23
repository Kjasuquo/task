run:
	go run main.go

mock:
	mockgen -source=database/dbInterface.go -destination=database/mocks/db_mock.go -package=mocks

