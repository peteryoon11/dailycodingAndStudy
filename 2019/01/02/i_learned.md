1. 자바 메모리에 따른 덤프를 뜨는 방법에 대해서 알아 봤다. 
    a. mat in eclipse 를 이용하는 부분을 알아 봤다. 
2. 자바 튜닝 하는 기준에 대해서 글을 읽어 봤다. 
jhat
jmeter

## 
• jmap: 힙 덤프나 히스토그램을 출력하는 프로그램
• jhat: 힙 덤프를 이용해 각 객체를 볼 수 있는 프로그램
• VisualVM: 실행되고 있는 JVM의 힙 내용을 볼 수 있는 프로그램
다음은 히스토그램의 일부다.




출처: <https://d2.naver.com/helloworld/1326256> 


Jmap -heap "pid"

Pid 는 
Ps -ef | grep "서비스 명으로 찾아서 입력"

MAT 로 힙덤프 를 떠서 분석을 하려면 

Jmap -dump:format=b,file=heap."pid".hpronf "pid"

로 파일을 만든 후에 로컬로 가져와서 eclipse 에서 mat 로 분석한다. 