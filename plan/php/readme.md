mac 에서 아파치 

https://gongdoll.tistory.com/14

sudo apachectl start

# 파일을 vi로 수정한다
sudo vi /etc/apache2/httpd.conf 

# 루트경로 설정값 변경
DocumentRoot "/Library/WebServer/Documents"
<Directory "/Library/WebServer/Documents"

sudo apachectl graceful

-----

https://hyangeun.tistory.com/8