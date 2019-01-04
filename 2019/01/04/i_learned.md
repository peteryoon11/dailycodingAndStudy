1. 스프링 부트의 application.yml 을 사용 할때 아래와 같은 profile 들이 존재 할때 
    application 뒤에 없는 부분이 있으면 기본 파일 즉 application.yml 를 불러온다. 
     application.yml application-dev.yml application-release.yml
     -Dspring.profiles.active=test -> application.yml  의 설정을 불러옴 