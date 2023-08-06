RELEASE_DB=postgresql://postgres:postgres@localhost:5432/release?sslmode=disable

release_migrate_up:
	migrate -path modules/release/db/migration -database ${RELEASE_DB} --verbose up

release_migrate_down:
	migrate -path modules/release/db/migration -database ${RELEASE_DB} --verbose up

sqlc:
	sqlc generate

.PHONY: release_migrate_up release_migrate_down sqlc