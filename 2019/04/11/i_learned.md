## 
1. centos 7 에서의 사용하는 포트 확인 및 사용하는 프로세스 확인 

사용하는 포트 

netstat -tulpn | grep LISTEN

https://flyyunha.tistory.com/entry/centos7-%EC%82%AC%EC%9A%A9%EC%A4%91%EC%9D%B8-%ED%8F%AC%ED%8A%B8-%ED%99%95%EC%9D%B8


https://medium.com/@jihasong/ubuntu-%EC%97%90%EC%84%9C-%EC%8B%A4%ED%96%89%EC%A4%91%EC%9D%B8-mysql%EC%9D%84-docker-container-%EB%82%B4%EB%B6%80%EC%9D%98-mysql%EB%A1%9C-%EB%A7%88%EC%9D%B4%EA%B7%B8%EB%A0%88%EC%9D%B4%EC%85%98-%ED%95%98%EA%B8%B0-e97d6243b0cf


docker run -it -p {port번호}:3306 --name {컨테이너이름} \
-e MYSQL_ROOT_PASSWORD={root_pw} \
-v /etc/localtime:/etc/localtime:ro \
-v /etc/timezone:/etc/timezone:ro \
-v {destination datadir}:{docker 내부 datadir} \
-d {image name}:{tag name}
///// 명령어 예시
docker run -it -p 3306:3306 --name mysql_server \
-e MYSQL_ROOT_PASSWORD=Password123! \
-v /etc/localtime:/etc/localtime:ro \
-v /etc/timezone:/etc/timezone:ro \
-v /mnt/my_db_hdd/mysql/:/var/lib/mysql/ \
-d mysql:5.6