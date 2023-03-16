#!/bin/bash

## 设置并加载变量
source "./_shell/init.sh"

## 判断参数
desc=$1
if [ -z "${desc}" ]; then
  echo -e "\033[31m Err:需要同步说明 \033[0m"
  exit 1
fi
echo "git commit: ${desc}"

GitSet

git pull &&
  git add . &&
  git commit -m "${desc}" &&
  git push &&
  echo "同步完成"

GitSet &&
  exit 0
