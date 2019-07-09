http://blog.naver.com/PostView.nhn?blogId=alice_k106&logNo=220847310053

### 저장소 주소 설정
sinopia 는 nmpjs.org 의 저장소와 연결하여 필요한 패키지를 전달해 준다. 만일 nmpjs.org 사이트가 닫혀있어도 이미 전에 다운 받은 적이 있던 패키지들은 sinopia에 담겨 있으므로 여전히 개발자들은 대신해서 패키지를 다운 받을 수 있다. 그렇다면 이제는 내 컴퓨터의 npm 은 지금 만들어논 사설 저장소를 사용하도록 설정해보자.
$ npm set registry “http://npm-registry.company.com:4873"
$ npm config get registry

-> https://medium.com/@yongho.hwang/npm-%EC%82%AC%EC%84%A4-%EC%A0%80%EC%9E%A5%EC%86%8C-%EB%A7%8C%EB%93%A4%EA%B8%B0-%EC%84%9C%EB%A1%A0-%EC%9D%B4-%EB%82%B4%EC%9A%A9%EC%9D%80-%EB%A7%88%ED%81%AC%EB%8B%A4%EC%9A%B4-%ED%98%95%EC%8B%9D%EC%9C%BC%EB%A1%9C-%EC%9E%91%EC%84%B1%EB%90%98%EC%97%88%EA%B3%A0-%EB%A7%88%ED%81%AC%EB%8B%A4%EC%9A%B4-%EB%AC%B8%EB%B2%95-http-blog-baenlee-com-blog-2013-08-13-bf08ef748bd1