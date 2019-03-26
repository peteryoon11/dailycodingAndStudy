## useradd 에 대한 정리 

https://eunguru.tistory.com/88

https://freehoon.tistory.com/39




how-to-run-java-jar-application-with-systemd-on-linux


https://computingforgeeks.com/how-to-run-java-jar-application-with-systemd-on-linux/




## jar 파일을 만들때 java -jar 없이 바로 실행 가능한 *.jar 을 만들기 위해서는 

build.gradle 에 아래의 부분을 추가 해야 한다.


jar {
    version = "0.0.1"
}

bootJar {
	archiveName = "excuatlbeName.jar"
    launchScript()
}