#!/bin/bash
echo "==================     Help for psql    ========================="
echo "\l or \list                : shows all databases"
echo "\d                         : shows all tables, views and sequences"
echo "\d table_name              : describe table, view, sequence, or index"
echo "\c database_name           : connect to a database"
echo "\q                         : quit"
echo "\?                         : for more commands"
echo "====================    Extensions    ==========================="
echo "create extension pgcrypto; : installs cryptographic functions"
echo "====================    Some SQL    ============================="
echo "select gen_random_uuid();  : returns a random uuid (pgcrypto)"
echo "select version();          : return the server version"
echo "select current_date;       : returns the current date"
echo "================================================================="
docker exec -it postgres psql -U postgres