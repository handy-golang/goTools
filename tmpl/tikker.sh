# !/bin/bash

path={{.Path}}
cd ${path}
shellFile=${path}"/{{.FileName}}"

{{.ShellCont}}

rm -rf ${path}"/tikker.sh"
pm2 delete tikker
