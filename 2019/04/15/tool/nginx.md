CentOS7 에 NginX 최신버전 ( mainline ) 설치하기

출처: https://haru.kafra.kr/106 [어제도 오늘도 내일도 언제나 하루]

vim /etc/yum.repos.d/nginx.repo

아래를 추가한다.(centos7 기준)



[nginx]

name=nginx repo

baseurl=http://nginx.org/packages/mainline/centos/7/$basearch/

gpgcheck=0

enabled=1





:wq 로 저장하고 나와서

yum install nginx



그리고 설정내용은 nginx 서버별 설정 노트를 참고한다.nginx 서버별 설정

시작프로그램에 등록

systemctl enable nginx



재부팅시 정상적으로 재시작이 되는것을 확인하기 위해 재부팅해준다.



centos7 부터는 firewall-cmd 를 기본으로 사용하게 된다.

80번 포트와 443 포트를 열어주도록 하자.

firewall-cmd --permanent --zone=public --add-port=80/tcp

firewall-cmd --permanent --zone=public --add-port=443/tcp

firewall-cmd --reload

firewall-cmd --list-all

https://www.nginx.com/resources/wiki/start/topics/examples/full/

