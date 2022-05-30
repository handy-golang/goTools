# !/bin/bash

path="{{.Path}}"
cd ${path}
serverName="{{.FileName}}"
shellFile=${path}"/{{.FileName}}"

rm -rf ${shellFile}

echo "开始执行脚本"

{{.ShellCont}}

echo "脚本执行结束,清理进程"

pm2 delete ${serverName}
