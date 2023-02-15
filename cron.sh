#!/bin/bash

DAY=$(date +"%d")

function is_prime() {
    if [[ $1 -eq 2 ]] || [[ $1 -eq 3 ]]; then
        return 1  # prime
    fi

    if [[ $(($1 % 2)) -eq 0 ]] || [[ $(($1 % 3)) -eq 0 ]]; then
        return 0  # not a prime
    fi

    i=5; w=2

    while [[ $((i * i)) -le $1 ]]; do
        if [[ $(($1 % i)) -eq 0 ]]; then
            return 0  # not a prime
        fi
        i=$((i + w))
        w=$((6 - w))
    done

    return 1  # prime
}

# Exit if current day is a prime number
is_prime $DAY
if [[ $? -eq 1 ]]; then
  exit 0
fi


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
git commit -m "$(curl --silent 'https://whatthecommit.com/index.txt')"
git push origin master

exit 0
