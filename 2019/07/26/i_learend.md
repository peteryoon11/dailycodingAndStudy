http://vezi95.blogspot.com/2016/05/git-pull.html


git pull 시 아래와 같은 에러가 뜰때 위의 내용으로 사용 함

error: The following untracked working tree files would be overwritten by checkout:

이유는 기존에 stage 안된 파일이 생겨서 삭제 하거나 이동 하라고 한다.

다른 저장소에서 넘어오면 보통 덮어 씌우면 되니까. 

git add -A
git stash
git pull 

의 방법으로 해결 함.

 git 은 매번 쓰는 내용만 쓰다보니까 가끔 오류를 마주보면 이상하게 덮어 씌울까봐 조심하게 된다.



 https://blog.azulpintor.io/entry/PostgreSQL-null-%EC%B2%B4%ED%81%AC%ED%95%98%EB%8A%94-%EB%B0%A9%EB%B2%95-COALESCE-%ED%95%A8%EC%88%98


 postgresql 에서 사용하는 COALESCE 함수 에 대한 내용 
 다른 db 에서는 사용해 본적이 없어서 기록을 하고 기억해 둬야 겠다.