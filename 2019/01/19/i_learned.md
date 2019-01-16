
oralce 에서 기존의 primary key 로 되어 있는 컬럼에 추가로 primary key를 등록 하기 위한 쿼리 

ALTER TABLE "테이블 이름" DROP CONSTRAINT "primary key 이름 ";




drop index "기존의 uniq index 가 연결 되어 있다면 해당 uniq index"; -- 여기 추가.. uniq index 를 안지웠음 

CREATE UNIQUE INDEX "primary key 이름" ON "테이블 이름"("추가할 컬럼 이름", "추가할 컬럼 이름");


ALTER TABLE "테이블 이름" ADD ( CONSTRAINT "primary key 이름" PRIMARY KEY ("추가할 컬럼 이름", "추가할 컬럼 이름") );