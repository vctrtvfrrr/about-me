#!/bin/bash

BIN=./bin/readme

[[ ! -f $BIN || ! -f .env ]] && echo "App does not installed."

$BIN

git checkout master
git add README.md
git commit -m "$(curl --silent 'http://whatthecommit.com/index.txt')
git push origin master
