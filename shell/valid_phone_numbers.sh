#!/bin/sh

# 193. 有效电话号码

# 给定一个包含电话号码列表（一行一个电话号码）的文本文件 file.txt
# 写一个单行 bash 脚本输出所有有效的电话号码。

grep -E '^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$' phone.txt

grep -E '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$' phone.txt

awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' phone.txt