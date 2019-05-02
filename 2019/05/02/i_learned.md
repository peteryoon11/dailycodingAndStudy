# ruby on rails begin

mac 
ruby -v 
gem install rails 

rails -v 

server 띄우기 

cd project_name/bin 

ruby rails server


controller 추가 하기 

cd project_name/bin 
(X) ruby rails generate controller Welcome index

(O) rails generate controller Welcome index

rails route

rails generate model Article title:string text:text


5.6 컨트롤러에서 데이터를 저장하기
이제 ArticlesController에 돌아갑시다. 좀 전에 만들었던 Article 모델을 사용해서 create 액션을 수정해야 합니다. app/controllers/articles_controller.rb를 에디터로 열어서 create 액션을 다음과 같이 변경하세요.