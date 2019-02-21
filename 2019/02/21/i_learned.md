## JPA 에서의 save 와 saveAndFlush 에 대해서 

현재 코드에서는 save 를 주로 사용 
# 관련해서 어느 순간에 save 를 쓰고 saveAndFlush 를 쓰는 지도 확인해 보자. 
 saveAndFlush db에 내용이 바로 반영 
 save 는 해당 트랜젝션이 끌날때 반영 

참조 :
https://cnpnote.tistory.com/entry/SPRING-Spring-%EB%8D%B0%EC%9D%B4%ED%84%B0-jpa%EC%97%90%EC%84%9C-save%EC%99%80-saveAndFlush%EC%9D%98-%EC%B0%A8%EC%9D%B4%EC%A0%90
https://ramees.tistory.com/36




java class 마다 있는 serialVerisonUID 의 용도에 대해서 
https://javafactory.tistory.com/1388

eclipse 에서나 intellij 에서나 기본적으로 serailVersionUID를 생성해 주지만 명시적으로  주는게 더 안전하다고 함 

https://civan.tistory.com/113



/**
		 *  generated serialVersionUID 를 통해서 만들어진 uid 
		 *  default 로 하면 1L 로 성의 없게 만들어진다.
		 */
		private static final long serialVersionUID = -1187593977680314664L;