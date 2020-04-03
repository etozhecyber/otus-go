PgBasePass ?= "Password"

migrate:
	cat *.sql | psql -v basepass="'${PgBasePass}'"
