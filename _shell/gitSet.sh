#!/bin/bash

echo "设置大小写敏感"
git config core.ignorecase false

echo "设置忽略权限变更"
git config --global core.fileMode false
git config core.filemode false

echo "更改文件夹权限"
chmod -R 755 ./
