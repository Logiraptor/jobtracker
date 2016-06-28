#!/bin/bash

pg_dump -U jobtracker jobtracker > "`dirname $0`/structure.sql"
