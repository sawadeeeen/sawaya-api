#!/bin/bash

echo `docker-compose build --no-cache`
echo `docker-compose up -d`
echo `docker-compose exec api migrate -database 'mysql://root:root@tcp(db:3306)/ianBlog' -path ./sql up`
