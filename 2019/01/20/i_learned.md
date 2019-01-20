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
--------------------

clean code 를 읽으면서 

2장 의미있는 이름을 읽고 있다. 

info , data 라는 말은 쓰지 말고 

이름에 대해서 정확히 무슨 일을 하는 것인지 알려줄것 
조금 길어 져도 상관없다는 것 

메소드도 마찬가지 

