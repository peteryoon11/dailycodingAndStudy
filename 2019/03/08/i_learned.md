
## gradle.build  내부에 대한 설명 
찾고 있던 부분인데 

gradle repository local 위치

위의 검색어로 찾아냄 

http://kwonnam.pe.kr/wiki/gradle/dependencies


JPA model gen 

http://kwonnam.pe.kr/wiki/gradle/jpa_metamodel_generation

// 리스트로 만들어진 객체에서 하나의 키로 그 키가 반복 되는 만큼의 숫자 세는 메소드 

public Map<String,Integer> CheckFrequencyProcessy(List<TestEntity> listemp){

    	List<String> templist = new ArrayList<>();
    	Map<String, Integer> result=new HashMap<>();
    	for(TestEntity item: listemp) {
    		templist.add(item.getAccountID());
    	}
    	 
    	 Set<String> mySet = new HashSet<String>(templist);
    	 for(String s: mySet){

    	
    	  result.put(s, Collections.frequency(templist,s));
    	 }
    	
    	return result;
    }