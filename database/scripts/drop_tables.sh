#!/usr/bin/env bash
psql -c "DROP TABLE user_roles" bfdb
psql -c "DROP TABLE role" bfdb
psql -c "DROP TABLE post_part" bfdb
psql -c "DROP TYPE post_part_type" bfdb
psql -c "DROP TABLE post" bfdb
psql -c "DROP TABLE \"user\"" bfdb
