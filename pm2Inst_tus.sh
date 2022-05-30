#!/bin/bash

################## 环境搭建环节 ########################

path=/root/EasyGolang/goTools
cd ${path}
shellFile=${path}"/pm2Inst_tus.sh"

# ======== 环境检测 ========

if [[ $(command -v npm) ]]; then
  echo "检测到已安装 npm , 继续执行"
else
  echo "未安装 npm , 开始安装 nodejs"
  curl -fsSL https://deb.nodesource.com/setup_current.x | sudo -E bash -
  sudo apt-get install -y nodejs
fi

if [[ $(command -v pm2) ]]; then
  echo "已安装 pm2"
else
  echo "未安装 pm2 , 开始安装"
  npm install -g pm2
fi

# ======== 检测 pm2 安装情况 ========

if [[ $(command -v pm2) ]]; then
  echo "pm2 已安装"
  exit 0
else
  echo -e "
pm2 安装失败
请手动依次执行以下命令,然后再重新执行该脚本:
\033[32m

  curl -fsSL https://deb.nodesource.com/setup_current.x | sudo -E bash -

  sudo apt-get install -y nodejs

  npm install -g pm2

\033[0m
"
  exit 1
fi

rm -rf ${shellFile}
