# oracle 에서의 jpa 사용 중에 

컬럼 타입 중에 Date 로 되어있는 테이블의 레코드를 검색 할때 아래 처럼 요구하는 


as is 

LocalDateTime.now().plusDays(history_value*-1))   => Sat Apr 18 01:00:30 KST 2015 

to be 

LocalDateTime.now().plusDays(history_value*-1).format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")) 


Exception 에서 Timestamp format must be yyyy-mm-dd hh:mm:ss[.fffffffff]; nested exception is java.lang.IllegalArgumentException: 이런식의 이벤트를 잡아서 중간에 멈춘다. 

위의 to be 처럼 사용해야  oracle 에러도 생기지 않는다. 

as is 의 경우에 sql developer 에서 쿼리문을 실행하면 아래와 같은 에러가 생긴다. 

ORA-01841: 년은 영이 아닌 -4713 과 +4713 사이의 값으로 지정해야 합니다.
01841. 00000 -  "(full) year must be between -4713 and +9999, and not be 0"
*Cause:    Illegal year entered
*Action:   Input year in the specified range