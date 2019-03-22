## 



EntitySpec 을 만들때 root.get 에 들어가는 건 아래처럼 문자열로 넣으면 안되는 거 같다. 
기존에 있던 것처럼 entitiy 별로 generated 된 

new Specification<TestEntity>() {

			@Override
            public Predicate toPredicate(Root<TestEntity> root, CriteriaQuery<?> query, CriteriaBuilder cb) {
                
                return cb.and( cb.lessThan(root.get("CREATED"), timePoint), cb.equal(root.get("PP_ID"), spid));
            }
        };









## 일단 아래 내용은 적용이 안됨.. spring boot 2.0 gradle 5.0 정확히는  annotationProcessor 를 어디에 넣어야 할지 알수 없음 

그에 따라서 gralde 5.0 에서 jpamodelgen 을 사용하기로 함  아래는 해당 내용에 대한 stackoverflow 와 내용 



https://stackoverflow.com/questions/54218556/how-to-generate-jpa-metamodel-with-gradle-5-x

you can just remove the plugin for the jpa modelgen and just use

annotationProcessor('org.hibernate:hibernate-jpamodelgen:<version>')
Addtionally i use those settings to configure where the generated code should live.

tasks.withType(JavaCompile) {
  options.annotationProcessorGeneratedSourcesDirectory = file("src/generated/java")
}

compileJava.options.compilerArgs += [
]

sourceSets {
    generated {
        java {
            srcDirs = ['src/generated/java']
        }
    }
}


위와 같은 검색어로 찾은 사이트 인데 딱히... 
http://kwonnam.pe.kr/wiki/gradle/jpa_metamodel_generation