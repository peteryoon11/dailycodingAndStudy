# centos 7 에서의 
1. mariadb10.3 설치 
2. root 비번 설정 
3. 유저 생성 
4. 유저에게 접속 할 ip 설정(대역 설정)과 스키마에 대한 권한 부여 
5. 사용할 포트 열기 



1. mariadb10.3 설치 

Centos 7 MariaDB 10.3 설치
https://blog.kerus.net/1022/install-yum-mariadb-10-3-on-centos7

# YUM REPO에 MARIADB 추가
sudo vim /etc/yum.repos.d/MariaDB.repo

[mariadb]
name = MariaDB
baseurl = http://yum.mariadb.org/10.3/centos7-amd64
gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB
gpgcheck=1

# 설치

sudo yum install MariaDB-client MariaDB-server


# 실행 및 서비스등록

systemctl start mariadb
systemctl enable mariadb


# MYSQL ROOT 계정접속

sudo mysql


2. root 비번 설정 

MySQL/MariaDB 초기 설치 후 계정 패스워드 초기화 (update root password)
https://sarc.io/index.php/mariadb/931-mysql-mariadb-update-root-password

#. 개요

최초 설치 시에는 root 계정 패스워드가 설정되어 있지 않다. root 계정의 패스워드를 설정하는 방법이다.

 

#. 방법

2-1. mysql 접속

패스워드가 설정되어 있지 않은 상태이므로 그냥 접속할 수 있다.

mysql -u root

2-2. mysql DB 선택

use mysql;

2-3. 계정 패스워드 초기화

update user set password=password('mypassword') where user='root';
flush pribvileges;




3. 유저 생성 및 스키마 생성 
https://sehoonoverflow.tistory.com/6


# 유저 생성 

create user '계정아이디'@'접속위치' identified by '패스워드';

ex. create user 'user1'@'%' identified by 'user!@#$';

# 스키마 생성 

CREATE DATABASE study_db default CHARACTER SET UTF8;



4. 유저에게 접속 할 ip 설정(대역 설정)과 스키마에 대한 권한 부여 


DB 권한 부여하기
$ grant all privileges on *.* to '사용자'@'localhost';
$ grant all privileges on DB이름.* to '사용자'@'localhost';


# 그외 알아 둘 부분 

권한 확인
# show grants for 'user1'@'접속위치';

계정 삭제
# drop user '계정아이디'@'접속위치';
ex. drop user 'user1'@'%';

권한 삭제
# revoke all on DB이름.테이블 FROM '계정아이디'@'접속위치';

출처: https://sehoonoverflow.tistory.com/6 [세훈오버플로우]

5. 사용할 포트 열기 

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
