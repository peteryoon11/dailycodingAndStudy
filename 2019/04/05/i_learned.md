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




CREATE DATABASE study_db default CHARACTER SET UTF8;



DB 권한 부여하기
$ grant all privileges on *.* to '사용자'@'localhost';
$ grant all privileges on DB이름.* to '사용자'@'localhost';


