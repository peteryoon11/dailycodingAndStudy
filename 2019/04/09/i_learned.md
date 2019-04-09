	
    java에서의 디버그 할때의 중요성.. 2번을 사용하자. 너무 오랜만이라 헷갈려서 해맴...
1. logger.info(e.getStackTrace().toString());
2. logger.info(e.getMessage());

비슷한 사례를 찾았는데.. 안된다. 분명히 똑같이 맞췄는데.. 
gradle clean 
project clean 도 했는데..
https://matthew.kr/spring-boot-jpa%EC%97%90%EC%84%9C-update-%EC%BF%BC%EB%A6%AC-%EC%82%AC%EC%9A%A9%EC%8B%9C-xxx-is-not-mapped-%EC%97%90%EB%9F%AC-%EB%B0%9C%EC%83%9D%EC%8B%9C-%EB%8C%80%EC%B2%98%EB%B0%A9%EB%B2%95/




public interface ABCEntityRepository extends
		JpaRepository<ABCEntity, ABCId>,
		JpaSpecificationExecutor<ABCEntity> {

@Transactional
	@Modifying
	@Query("delete from ABC  where ID =:id and T_ID = :t_id")
	void deleteAllByIdAndtId(@Param("id")String id, @Param("t_id") String t_id);
        
        
        
        }

위와 같은 경우에 

as is 
@Query("delete from ABC  where ID =:id and T_ID = :t_id")


to be 

@Query("delete from ABCEntity  where ID =:id and T_ID = :t_id")

로 매핑이 되는 경우가 있는데...

다른 곳에서는 똑같은것이 없는 경우가 있다. 

table 이름 자리에 실제 테이블 이름을 적었는데... 
여기는 왜...Entity 의 이름을 적는거지....


# gunzip 사용하기 
https://whackur.tistory.com/112
