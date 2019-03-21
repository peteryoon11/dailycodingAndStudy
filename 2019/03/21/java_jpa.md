https://happygrammer.tistory.com/149







예) 루트 엔티티 선택
Criteria API를 이용하려면 CriteriaBuilder를 이용합니다. 다음은 CriteriaBuilder를 이용해 루트 엔티티를 선택하는 쿼리입니다.

CriteriaBuilder builder = entityManager.getCriteriaBuilder();
​
CriteriaQuery<Person> criteria = builder.createQuery( Person.class );
Root<Person> root = criteria.from( Person.class );
criteria.select( root );
criteria.where( builder.equal( root.get( Person_.name ), "John Doe" ) );
​
List<Person> persons = entityManager.createQuery( criteria ).getResultList();

