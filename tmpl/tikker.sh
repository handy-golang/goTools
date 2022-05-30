# !/bin/bash

path="{{.Path}}"
cd ${path}
serverName="{{.FileName}}"
shellFile=${path}"/{{.FileName}}"
logFile="{{.LogPath}}"

mEcho() {
  echo $@ >>${logFile}
}

rm -rf ${shellFile}

mEcho "==== 开始执行脚本 ===="

{{.ShellCont}}

mEcho "==== 脚本执行结束 ===="

pm2 delete ${serverName}
