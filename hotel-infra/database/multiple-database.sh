#!/bin/bash

# Exit immediately if a command exits with a non zero status
set -e
# Treat unset variables as an error when substituting
set -u

# Function to create a database
function create_databases() {
    database=$1
    echo "Creating database '$database'"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
      CREATE DATABASE $database;
      GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL

 if [ "$database" == "hotel_ums" ]; then
      echo "Creating ENUM in database '$database'"
      psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$database" <<-EOSQL
        CREATE TYPE user_role AS ENUM ('admin', 'guest');
EOSQL
    fi
}

# Check if POSTGRES_MULTIBLE_DATABASES is set
if [ -n "$POSTGRES_MULTIBLE_DATABASES" ]; then
  echo "Multiple database creation requested: $POSTGRES_MULTIBLE_DATABASES"
  for db in $(echo $POSTGRES_MULTIBLE_DATABASES | tr ',' ' '); do
    database=$(echo $db | awk -F":" '{print $1}' | tr '-' '_')  # Ganti tanda hubung dengan underscore
    echo "Creating database: $database"
    create_databases $database
  done
  echo "Multiple databases created!"
fi