#!/bin/sh

# 192. 统计词频
# 写一个 bash 脚本以统计一个文本文件 words.txt 中每个单词出现的频率。
cat words.txt | xargs -n 1 | sort | uniq -c | sort -rn | awk '{print $2,$1}'

