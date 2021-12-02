#!/bin/sh
# wait-for-postgres.sh

set -e

shift

until PGPASSWORD=1234 psql -h "db" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec "$@"