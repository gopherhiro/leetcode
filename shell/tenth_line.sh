#!/bin/sh

# 195. 第十行
# 给定一个文本文件 file.txt，请只打印这个文件中的第十行。
awk 'NR == 10' file.txt # NR在awk中指行号

sed -n 10p file.txt # -n表示只输出匹配行，p表示Print

tail -n+10 file.txt|head -n 1 # tail -n +10表示从第10行开始输出