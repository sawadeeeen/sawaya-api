#!/bin/bash

echo `docker-compose -f ./deployments/docker-compose.yml build --no-cache`
echo `docker-compose -f ./deployments/docker-compose.yml up -d`
echo `docker-compose -f ./deployments/docker-compose.yml exec api migrate -database 'mysql://root:root@tcp(db:3306)/ianBlog' -path ./sql up`
