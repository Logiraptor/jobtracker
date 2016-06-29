#!/bin/bash

set -xe

goose up

STRUCTURE="`dirname $0`/db/structure.sql"

pg_dump -s -U jobtracker jobtracker > "$STRUCTURE"


dropdb --if-exists -U jobtracker jobtracker-test
createdb -U jobtracker jobtracker-test
psql -U jobtracker jobtracker-test < "$STRUCTURE"