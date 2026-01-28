#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -U "$DB_USER" -d "$DB_NAME" -f /docker-entrypoint-initdb.d/database.sql;
