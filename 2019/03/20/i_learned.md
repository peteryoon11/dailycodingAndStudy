## jpa 에서의 오래된 히스토리 레코드 지우는 것과 관련 된 글 

https://jojoldu.tistory.com/235
-> 하나씩 지울때랑 여러개로 지울때 
여러개로 지울때는 select 해와서 해당 엔티티를 하나씩 날리는 방식으로 지운다. 
자세한 내용은 위에 있다.

하나의 사례로는 pk 가 걸려있는 하나를 지웠더니 그게 하나인것을 확인하고 해당하는 pk 가 걸려있는 3개의 필드에 같이 where 를 적용해서 지우더라..




스프링 데이터 JPA 레퍼런스 번역
http://arahansa.github.io/docs_spring/jpa.html


java LocalDate 와 LocalDateTime 을 구분하자... 같은 형식이면 LocalDateTime 으로 사용 하자. 



centos7 에 서비스 등록 하는 *.service 에 대해서 확인하다 보니 아래의 글이 나오더라. 

아래의 2개의 값에 대해서 확인하자. / 추가 스터디 필요 
[Service]
LimitNOFILE=65535
LimitNPROC=8192

https://fredrikaverpil.github.io/2016/04/27/systemd-and-resource-limits/


https://www.freedesktop.org/software/systemd/man/systemd.exec.html

https://victorydntmd.tistory.com/215