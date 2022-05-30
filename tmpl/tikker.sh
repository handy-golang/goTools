# !/bin/bash

path={{.Path}}
cd ${path}

{{.ShellCont}}

rm -rf ${path}"/tikker.sh"
pm2 delete tikker
