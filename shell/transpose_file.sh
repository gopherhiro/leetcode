#!/bin/sh

# 194. 转置文件
# 给定一个文件 file.txt，转置它的内容。(你可以假设每行列数相同，并且每个字段由 ' ' 分隔)。

# 获取第一行，然后用wc来获取列数
cols=`head -1 f.txt | wc -w`

# 循环所有的列
for (( i = 1; i <= $cols; i++ )); 
do
    # 使用 awk 依次输出文件的每一列，然后使用 xargs 转置即可。
    awk -v col=$i '{print $col}' f.txt | xargs
done