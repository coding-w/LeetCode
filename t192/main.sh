#!/bin/bash

# 统计词频
cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -nr | awk '{print $2, $1}'

# cat words.txt：读取文件内容。
# tr -s ' ' '\n'：将空格替换为换行符，确保每个单词占一行。
# sort：对所有单词按字典序排序。
# uniq -c：统计每个单词的出现次数，并输出格式为：次数 单词。
# sort -nr：根据次数降序排列。
# awk '{print $2, $1}'：调整输出格式为：单词 次数。
