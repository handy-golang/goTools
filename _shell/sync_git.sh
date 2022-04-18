#!/bin/bash

desc=$1

if [ -a ${desc} ]; then
  desc="exit-push"
fi

git pull &&
  git add . &&
  git commit -m "${desc}" &&
  git push &&
  echo "同步完成"
exit
