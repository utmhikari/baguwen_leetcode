#192. 统计词频
#写一个 bash 脚本以统计一个文本文件words.txt中每个单词出现的频率。
#
#为了简单起见，你可以假设：
#
#words.txt只包括小写字母和' '。
#每个单词只由小写字母组成。
#单词间由一个或多个空格字符分隔。

cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{print $2, $1}'
