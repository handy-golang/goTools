#!/bin/bash

################## 环境搭建环节 ########################

echo "
======== 环境检测 ========
"

if [[ $(command -v npm) ]]; then
  echo "检测到已安装 npm , 继续执行"
else
  echo "未安装 npm , 开始安装 nodejs"
  curl -fsSL https://deb.nodesource.com/setup_current.x | sudo -E bash -
  sudo apt-get install -y nodejs
fi

if [[ $(command -v pm2) ]]; then
  echo "已安装 pm2 , 继续执行"
else
  echo "未安装 pm2 , 开始安装"
  npm install -g pm2
fi
