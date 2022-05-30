# !/bin/bash

path="{{.Path}}"
cd ${path}
serverName="{{.FileName}}"
shellFile=${path}"/{{.FileName}}"
logFile="{{.LogPath}}"

rm -rf ${shellFile}

echo "开始执行脚本" >>${logFile}

{{.ShellCont}}

echo "脚本执行结束,清理进程" >>${logFile}

pm2 delete ${serverName}
