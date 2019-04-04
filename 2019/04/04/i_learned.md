# docker mysql mariadb 


# 안맞음  , mariadb 10.3 버전 으로 깔아서 해결 변경 해야 하는 것들이 여러가지 있음 
Can't connect to local MySQL server through socket '/var/lib/mysql/mysql.sock' (2)

출처: https://algo79.tistory.com/entry/MySQL-에러 [이제부터 시작!! 어디 한번해볼까?]


https://algo79.tistory.com/entry/MySQL-%EC%97%90%EB%9F%AC





# 안맞음
[문제해결 일지] systemctl start mariadb 실행시 Failed to issue method call: No such file or directory. 에러

https://conory.com/blog/43046





https://blog.kerus.net/1022/install-yum-mariadb-10-3-on-centos7



# 맞음 
mariadb 처음 깔고 나서 root 비번 설정 

적용 

flush privileges;

https://3unkong.tistory.com/17



마리아디비 
https://digndig.kr/mysql/793/




CentOS 7에서 포트 열기


# 현재 열려있는 포트를 확인
netstat -tulpn | grep LISTEN

# CentOS6까지만 동작했던 명령입니다.
service iptables status

# CentOS7에서 방화벽 iptables 현황을 볼 수 있습니다.
iptables -L --line

# 외부 테스트를 위해서 웹서버를 띄웁니다.
python -m SimpleHTTPServer 8000

# 포트가 외부에서 접속되지 않는다면 포트를 방화벽에 추가합니다.
sudo firewall-cmd --zone=public --add-port=8000/tcp --permanent

# 방화벽을 리로드합니다.
sudo firewall-cmd --reload