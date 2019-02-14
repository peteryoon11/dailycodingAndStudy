##  centos 7 temp dir 
# centos 7 에는 systemd-tmpfiles-clean.service  라고 하는 temp 파일을 주기적으로 정리하는 서비스가 있다. 

아래에서 지우는 주기나 위치 들을 조정 가능 한듯 

cat /usr/lib/tmpfiles.d/*.conf

아래 참고 사이트는 RHEL7 기준 인듯 그런데 centos7 이나 RHEL7 이나 모두 같은 거로 

알고있는데 다르다..

원인은 좀 더 찾아봐야 겠다.

참고 사이트 

https://developers.redhat.com/blog/2016/09/20/managing-temporary-files-with-systemd-tmpfiles-on-rhel7/

-----------------


java 의 설치 정보 들 
temp ? heap 메모리 , os info 
java -XshowSettings


java option 


java -Djava.io.tmpdir=/path/to/tmpdir

-----------
`
jinfo [PID]
`