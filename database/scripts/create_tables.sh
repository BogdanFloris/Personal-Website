#!/usr/bin/env bash
psql -f ./database/tables/user.sql bfdb
psql -f ./database/tables/role.sql bfdb
psql -f ./database/tables/user_roles.sql bfdb
psql -f ./database/tables/post.sql bfdb
psql -f ./database/tables/post_part.sql bfdb
