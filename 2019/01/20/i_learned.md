https://stackoverflow.com/questions/21529437/how-to-do-natural-sort-output-of-uniq-c-in-descending-acsending-order-unix




aaa dfdf
aaa dfdf
bbbd dfdf
ccc
ddd
ddd
aaa


sort test.txt | uniq -c | sort -nr

      2 ddd
      2 aaa dfdf
      1 ccc
      1 bbbd dfdf
      1 aaa

grep -v DDD