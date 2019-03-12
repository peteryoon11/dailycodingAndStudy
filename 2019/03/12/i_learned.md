# json 파일의 크기 한계에 대해서 알아 보았다.

https://knight76.tistory.com/entry/30021323262
https://www.ibm.com/support/knowledgecenter/ko/SS9H2Y_7.6.0/com.ibm.dp.doc/json_parserlimits.html

post 에서의 content length 에 보내는 용량을 명시 해야 하고. 

그것이 아니더라도 tomcat 의 경우에는  아래 처럼 default 용량 제한이 생긴다고 한다. 
그렇지만 아래의 용량을 풀어버리면 BTOC 의 경우  ddos 공격에 취약해 진다는 의견도 있다. 


데이터의 용량이 2기가를 넘는다면 에러를 발생하게 되는데 톰캣의 경우는 아래와
같이 Server.xml에서 maxPostSize를 5기가 정도를 늘려주면 된다.
설정이 없다면 2097152(2 megabytes)가 디폴트이다.
http://blog.moramcnt.com/?p=980
