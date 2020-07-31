#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
BIN=./bin/readme

cd $DIR

if [[ ! -f $BIN || ! -f $DIR/.env ]] ; then
    echo "App does not installed."
    exit 1
fi

$BIN

git checkout master
git add README.md
git commit -m "$(curl --silent 'http://whatthecommit.com/index.txt')"
git push origin master

exit 0
