#!/bin/bash

lsof -ti :3000 | xargs kill

npm run build 

if [ $? -ne 0 ]; then
  echo "npm run build failed. Aborting further steps."
  exit $?
fi

cp -r dist ./go 

cd ./go 

air -c .air.toml

lsof -ti :3000 | xargs kill
