## 부하 쿼리와 관련된 분석 및 개선 방안에 대해서 

# avr 200~300ms 인데 사용자의 모든 정보를 가지고 와서 offset 를 걸어서 보여주는데..

# 캐쉬를 해야 할까? 


https://cwondev.tistory.com/15

https://docs.oracle.com/javase/8/docs/api/java/util/Collection.html

private static <E, V> List<V> mapMakeTest(Collection<E> valueTest, Function<? super E, V> valueFunction) {
    
        return valueTest.stream()
                .map(valueFunction)
                .collect(Collectors.toList());
    }
    