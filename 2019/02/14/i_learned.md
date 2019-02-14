##  centos 7 temp dir 
# centos 7 에는 systemd-tmpfiles-clean.service  라고 하는 temp 파일을 주기적으로 정리하는 서비스가 있다. 

아래에서 지우는 주기나 위치 들을 조정 가능 한듯 

cat /usr/lib/tmpfiles.d/*.conf


참고 사이트 

https://developers.redhat.com/blog/2016/09/20/managing-temporary-files-with-systemd-tmpfiles-on-rhel7/
