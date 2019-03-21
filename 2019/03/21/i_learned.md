## linux sort 


http://hochulshin.com/linux-sort-reverse-order/
du | sort -n 
숫자 오름차순 순서대로 


Linux 유용한 bash 명령들

http://hochulshin.com/linux-useful-bash-commands/



압축 파일들이 있는 폴더에서 크기 역순으로 정렬해서 가장 큰 파일 열기

ls -al | sort -k 5 -nr | head -1 | cut -d " " -f 10 | xargs zmore



ls -al: 폴더 파일들 목록 출력
sort -k 5 -nr: 5번째 컬럼(여기서는 파일 크기)을 역순으로 정렬
head -1 : 첫번째 아이템 만
cut -d “ “ -f 10: 10번째 컬럼만 잘라내기 (여기서는 파일 이름)
xargs zmore: 앞의 파이프 입력을 아큐먼트로 zmore 실행