#!/bin/bash
function wait_for_server() {
  until `docker exec postgres pg_isready -p 5432 -h localhost -q`; do
    sleep 1
  done
}
function install_dvdrental() {
unzip dvdrental/dvdrental.zip -d dvdrental
docker exec postgres pg_restore -U postgres -d postgres /dvdrental/dvdrental.tar
rm -f dvdrental/dvdrental.tar
}
docker rm -f $(docker ps -aq)
docker-compose up -d
echo "Waiting on postgres server..."
wait_for_server
install_dvdrental
./psql-cli.sh