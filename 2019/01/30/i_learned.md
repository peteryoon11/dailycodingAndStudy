oracle redo log 

관련 내용 

crontab 에 등록할 shell script 를 만들었다 .

파일 오픈 갯수를 확인 하는 부분 


########################3
`
#!/bin/bash
 
# filename log-record-total.sh
 
NOW=$(date '+%G-%m-%e %H:%M:%S')
 
DAY=$(date  '+%Y%m%d')
 
LOG="/home/testuser/sh-test/testlog-total-$DAY.log"
 
content=$(cat /proc/sys/fs/file-nr)
 
echo "[$NOW] $content " >>$LOG
`

# 매분 마다 bizcpn 의 파일 사용 갯수를 확인 해서 로그로 남김
* * * * * /home/testuser/sh-test/log-record-total.sh
