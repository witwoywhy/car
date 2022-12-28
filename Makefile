unittest:
	GO_TEST_STAGE=mock go test ./... -v

integration:
	docker compose up -d
	sleep 1
	GO_TEST_STAGE=integration go test ./... -v
	docker compose down

clean:
	docker compose down