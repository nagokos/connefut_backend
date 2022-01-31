DC = docker-compose
DC_WEB = $(DC) exec web
DC_DB = $(DC) exec db
POSTGRESQL_URL = postgres://root:password@db:5432/connefut_db?sslmode=disable

package_add:
	@$(eval PACKAGE_NAME := $(shell read -p "add packages: " NAME; echo $$NAME))

package_name:
	@$(eval PACKAGE_NAME := $(shell read -p "packages name: " NAME; echo $$NAME))

schema_name:
	@$(eval SCHEMA_NAME := $(shell read -p "schema name: " NAME; echo $$NAME))

seq_name:
	@$(eval SEQ_NAME := $(shell read -p "seq name: " NAME; echo $$NAME))

force_version:
	@$(eval FORCE_VERSION := $(shell read -p "force version: " NAME; echo $$NAME))

# ** Go Command **
go_run: package_name
	${DC_WEB} go run ${PACKAGE_NAME}

go_get: package_add
	${DC_WEB} go get -d ${PACKAGE_NAME}

go_install: package_add
	${DC_WEB} go install ${PACKAGE_NAME}

# GraphQL Server Start
start_server:
	${DC_WEB} go run graph/server/server.go

mod_tidy:
	${DC_WEB} go mod tidy

# entとgqlgenの両方更新される
generate:
	${DC_WEB} go generate ./...

# ** DB関係 **
init_schema: schema_name
	${DC_WEB} go get -d entgo.io/ent/cmd/internal/printer@v0.9.2-0.20210821141344-368a8f7a2e9a && \
	          go get -d entgo.io/ent/cmd/ent@v0.9.2-0.20210821141344-368a8f7a2e9a && \
	          go run entgo.io/ent/cmd/ent init ${SCHEMA_NAME} 

install_migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# マイグレーションファイル作成
# SEQ_NAMEはcreate_users_tableの部分
init_migration: seq_name
  ${DC_WEB} migrate create -ext sql -dir db/migrations -seq ${SEQ_NAME}

# SQLのみを発行
# make generate後に使用
sql_migration:
	${DC_WEB} go run db/migrations/migrate.go

# データベースに反映
migrate_up:
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} up 1

migrate_down:
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} down 1

migrate_force: force_version
	${DC_WEB} migrate -path db/migrations -database ${POSTGRESQL_URL} force ${FORCE_VERSION}

create_initial_data:
	${DC_WEB} go run db/initial_data/data.go

db_attach:
	${DC_DB} bash