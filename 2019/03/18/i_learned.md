# Quartz에서 CronTrigger 사용하여 오래된 테이블 레코드를 지우기 


1. 테이블에 히스토리 저장 
2. 조건을 주고 해당 조건에 맞게 삭제
3. 삭제한 내역에 대해서 근거? 


아래 내용을 보고 따라 하자.
https://examples.javacodegeeks.com/enterprise-java/quartz/quartz-scheduler-tutorial/


transionManager 관련 내용 
현재 테스트 중인 부분에서 계속 transionmanager 가 없다고 에러를 내고 있음 
https://supawer0728.github.io/2018/03/22/spring-multi-transaction/