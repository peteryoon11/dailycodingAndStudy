
`
public interface CoinRepository extends JpaRepository<CoinEntity, String>, JpaSpecificationExecutor<CoinEntity> {
	
	
	
	
//	use_tp
	List<CoinEntity> findByCREATED(String CREATED);
	
    //List<CoinEntity> findByuse_tp(String use_tp);
	
	
    
   
}`

## JPA 사용 중에 entity 와 연결된 Repoisty를 만드는데 테이블의 컬럼에 _ 가 들어간 경우에는 findBy 에 넣어도 자꾸 에러가 생긴다. 
* IDE(Eclipse)  에러 라는 말이 있어서 컴파일도 해봤지만 동일 한걸로 봐서 정책인듯 하다. 그런데 관련해서 문서를 찾아봐도.. 그런 내용이 없는지. 못찾는건지... 