## useradd 에 대한 정리 

https://eunguru.tistory.com/88

https://freehoon.tistory.com/39




how-to-run-java-jar-application-with-systemd-on-linux


https://computingforgeeks.com/how-to-run-java-jar-application-with-systemd-on-linux/




## jar 파일을 만들때 java -jar 없이 바로 실행 가능한 *.jar 을 만들기 위해서는 

build.gradle 에 아래의 부분을 추가 해야 한다. / gradle 4.3 이상에서 동작 확인 


jar {
    version = "0.0.1"
}

bootJar {
	archiveName = "excuatlbeName.jar"
    launchScript()
}


------------- 

조금 더 낮은 gralde 에서는 

springBoot  {
    buildInfo()
    executable = true
}


https://theuphill.tistory.com/12


Spring Boot를 $ java -jar 명령어를 이용해 실행할 수도 있지만 spring boot 1.3.0 버전에서 추가된 기능을 이용해 별도 스크립트 없이 리눅스 환경에서 서비스로 등록해서 실행, 관리할 수 있다. 스프링 레퍼런스에서는 이런 형태를 완전히 실행 가능한 jar (fully executable jar)로 표현하고 있다.

실행 가능한 jar를 만들기 위해서는 maven에 아래 plugin을 추가해야 한다.



출처: https://theuphill.tistory.com/12 [오르막길을 올라가는 개발자]




------------------------


https://www.snoopybox.co.kr/1613

윈도우에서 작성한 Shell 스크립트 리눅스에서 에러나는 경우

에 대해서 찾아보면 위와 같은 경우도 나오지만 실제로는 그냥 윈도우에 있던 것을 복사해서 bash 창에서 복사하면 문자열만 복사 되는 듯 하다.