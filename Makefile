
ENV_FILE ?= .env

-include $(ENV_FILE)
export

db_host = localhost
ifdef PGHOST
	db_host = $(PGHOST)
endif

db_user = 5432
ifdef PGPORT
	db_port = $(PGPORT)
endif

db_user = tim
ifdef PGUSER
	db_user = $(PGUSER)
endif

db_pass =
ifdef PGPASSWORD
	db_pass := "\:$(PGPASSWORD)"
endif

db_name = app_db
ifdef PGDATABASE
	db_name := $(PGDATABASE)
endif

database_url = postgresql://$(db_user)$(db_pass)@$(db_host):5432/$(db_name)
ifdef DATABASE_URL
	database_url = $(DATABASE_URL)
endif


.PHONY: test build air hair view reset riverui

dump:
	@echo "HOST =" $(HOST)
	@echo "PORT =" $(PORT)
	@echo "PGDATABASE =" $(PGDATABASE)
	@echo "PGHOST =" $(PGHOST)
	@echo "PGPORT =" $(PGPORT)
	@echo "PGUSER =" $(PGUSER)
	@echo "PGPASSWORD =" $(PGPASSWORD)
	@echo "DATABASE_URL =" $(DATABASE_URL)
	@echo "database_url =" $(database_url)
	@echo "db_name =" $(db_name)


test:
	go test ./...

build:
	go build ./cmd/server

ci: build test

run:
	./server -log=json | humanlog --truncate=false --skip-unchanged=false

air:
	air

hair:
	air -- -log=json | humanlog --truncate=false --skip-unchanged=false

view:
	open $(HOST)

reset: reset-storage db-reset

hard-reset: reset-storage db-hard-reset

reset-storage:
	rm -f storage/public/originals/*
	rm -f storage/public/previews/*

gen:
	go generate ./...

db-clean:
	psql -c "select pg_terminate_backend(pid) from pg_stat_activity where pid <> pg_backend_pid() and datname = '$(db_name)';"
	-dropdb $(db_name)


db-hard-reset:
	-psql -c "select pg_terminate_backend(pid) from pg_stat_activity where pid <> pg_backend_pid() and datname = '$(db_name)';"
	-dropdb $(db_name)
	createdb $(db_name)
	river migrate-up --database-url="$(database_url)"
	cd migrations && \
	tern migrate && \
	tern code install code

db-river-migrate:
	river migrate-up --database-url="$(database_url)"


db-reset:
	psql $(db_name) < scripts/sql/reset.sql
	river migrate-up --database-url="$(database_url)"
	cd migrations && \
	tern migrate && \
	tern code install code

db-seed:
	psql $(db_name) < scripts/sql/seed.sql


tailwind-app:
	cd app; npx tailwindcss -i styles.css -o ../static/public/css/app.css --watch

tailwind-admin:
	cd admin; npx tailwindcss -i styles.css -o ../static/public/css/admin.css --watch

static/public/css/app.css: app/tailwind.config.js app/styles.css app/*.templ
	cd app; npx tailwindcss -i styles.css -o ../static/public/css/app.css --jit --minify

static/public/css/admin.css: admin/tailwind.config.js admin/styles.css admin/*.templ 
	cd admin; npx tailwindcss -i styles.css -o ../static/public/css/admin.css --jit --minify


stripe-listen:
	stripe --api-key $(STRIPE_API_KEY) listen --forward-to $(HOST)/stripe/webhook/

riverui:
	DATABASE_URL=$(database_url) PORT=8888 ./riverui

docker-run:
	docker run --env PGHOST=host.docker.internal --env PGUSER=tim --env PGDATABASE=app_db -p 8080:8080 tstack_app


## Deloyment
deploy: docker-build docker-push

docker-build:
	docker buildx build --platform linux/amd64 -t tstack_app:latest .

