
oralce 에서 기존의 primary key 로 되어 있는 컬럼에 추가로 primary key를 등록 하기 위한 쿼리 

ALTER TABLE "테이블 이름" DROP CONSTRAINT "primary key 이름 ";




drop index "기존의 uniq index 가 연결 되어 있다면 해당 uniq index"; -- 여기 추가.. uniq index 를 안지웠음 

CREATE UNIQUE INDEX "primary key 이름" ON "테이블 이름"("추가할 컬럼 이름", "추가할 컬럼 이름");


ALTER TABLE "테이블 이름" ADD ( CONSTRAINT "primary key 이름" PRIMARY KEY ("추가할 컬럼 이름", "추가할 컬럼 이름") );


java 8 에서 사용하는 특징들에 대해서 좀 더 찾아 보자.

https://lng1982.tistory.com/240


Wildcard Generic Type

?는 알 수 없는 타입

<?> : 모든 객체 자료형, 내부적으로는 Object로 인식
<? super 객체자료형> : 명시된 객체자료형의 상위 객체, 내부적으로는 Object로 인식
<? extends 객체자료형> : 명시된 객체 자료형을 상속한 하위객체, 내부적으로는 명시된 객체 자료형으로 인식


---------------------


파이썬으로 메일 보내기 예제 
http://pythonstudy.xyz/python/article/508-%EB%A9%94%EC%9D%BC-%EB%B3%B4%EB%82%B4%EA%B8%B0-SMTP


파이썬을 사용하여 HTML 이메일 보내기
https://stackoverrun.com/ko/q/12336929