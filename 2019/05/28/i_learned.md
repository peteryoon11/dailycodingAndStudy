## 영문으로 적는 코딩 테스트

https://medium.com/@joongwon/java-java%EC%9D%98-generics-604b562530b3

extends Comparable

연산자에 대한 

함수의 시그니처에 선언된 <T extends Comparable<T>> 은 “자신과 비교될 수 있는 모든 타입 T”라고 읽을 수 있다. 이렇게 설명되어 있는데 재귀적 타입 바운드가 존재하는 이유는 Java 언어의 문제점을 보완하기 위한 것이다.

Java는 연산자 오버 로딩을 지원하지 않기 때문에 short, int, double 등과 같은 primitive 타입에만 >와 같은 비교연산자를 사용할 수 있었다. 이 문제를 해결하기 위해 Comparable 인터페이스와 재귀적 타입 바운드를 활용한다. 예제 코드를 보면 이해에 도움이 될 것이다.

```
// 실행 안됨 클래스상태에서의 비교 블가
public static <T> int countGreaterThan(T[] anArray, T elem) {
    int count = 0;
    for (T e : anArray)
        if (e > elem)  // compiler error
            ++count;
    return count;
}
// 실행 가능 <T extends Comparable<T>> 선언되었기 때문에 
public static <T extends Comparable<T>> int countGreaterThan(T[] anArray, T elem) {
    int count = 0;
    for (T e : anArray)
        if (e.compareTo(elem) > 0)
            ++count;
    return count;
}
```