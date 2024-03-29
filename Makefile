include ./env/develop.env

DC = docker compose
DC_WEB = $(DC) exec web
DC_DB = $(DC) exec db
POSTGRESQL_URL = ${DB_URL}

dc_build:
	${DC} build

dc_up:
	${DC} up -d

dc_down:
	${DC} down

dc_ps:
	${DC} ps

package_add:
	@$(eval PACKAGE_NAME := $(shell read -p "add packages: " NAME; echo $$NAME))

package_name:
	@$(eval PACKAGE_NAME := $(shell read -p "packages name: " NAME; echo $$NAME))

schema_name:
	@$(eval SCHEMA_NAME := $(shell read -p "schema name: " NAME; echo $$NAME))

seq_name:
	@$(eval SEQ_NAME := $(shell read -p "seq name: " NAME; echo $$NAME))

version:
	@$(eval VERSION := $(shell read -p "version: " NAME; echo $$NAME))

# ** Go Command **
go_run: package_name
	${DC_WEB} go run ${PACKAGE_NAME}

go_get: package_add
	${DC_WEB} go get -d ${PACKAGE_NAME}

go_install: package_add
	${DC_WEB} go install ${PACKAGE_NAME}

lint:
	${DC_WEB} golangci-lint run

# GraphQL Server Start
start_server:
	${DC_WEB} go run graph/server/server.go

mod_tidy:
	${DC_WEB} go mod tidy

generate:
	${DC_WEB} go generate ./...

init_recruitments_data:
	${DC_WEB} go run db/test_data/recruitment/recruitment.go

init_users_data:
	${DC_WEB} go run db/test_data/user/user.go

init_follow_data:
	${DC_WEB} go run db/test_data/relationship/relationship.go

# ** DB関係 **
install_migrate:
	${DC_WEB} go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install_tbls:
	${DC_WEB} go install github.com/k1LoW/tbls@main

tbls:
	${DC_WEB} tbls doc -f

# マイグレーションファイル作成
# SEQ_NAMEはcreate_users_tableの部分
init_migration: seq_name
	${DC_WEB} migrate create -ext sql -dir db/migrations -seq ${SEQ_NAME}

# データベースに反映
migrate_up: version
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} up ${VERSION}

migrate_down: version
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} down ${VERSION}

migrate_reset:
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} down && ${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} up

migrate_force: version
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} force ${VERSION}

create_initial_data:
	${DC_WEB} go run db/initial_data/data.go

db_attach:
	${DC_DB} bash