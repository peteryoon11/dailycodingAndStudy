http://blog.web-uhee.com/archives/209


1
$ docker pull nginx
Nginx Docker 이미지를 제대로 Pull 받았는지, 다음의 명령어로 확인합니다.

1
$ docker images
목록에 Nginx repository 항목이 있으면 제대로 Pull 받은 것입니다. 이 Docker 이미지를 가지고 Docker 컨테이너를 생성하고 구동하겠습니다. 다음의 명령어로 간단히 생성해서 구동할 수 있습니다.

1
$ docker run --name uhee-proxy --detach nginx
별 오류가 발생하지 않았다면, 다음의 명령어로 구동 여부를 확인할 수 있습니다.

1
$ curl http://localhost
Welcome to Nginx html 파일의 소스가 출력된다면, Nginx Docker 컨테이너는 제대로 구동되고 있습니다. 다음의 명령어로 Docker 컨테이너의 목록과 이들의 상태를 확인할 수 있습니다.

1
$ docker ps -a
이 Docker 컨테이너로 구동하는 가상 머신은, 호스트 OS 와 파일 시스템이 논리적으로 고립되어 있습니다. 가상 머신을 사용하는 이유이기도 하지만 그 때문에 Docker 컨테이너에 있는 데이터를 관리하는게 여간 불편한 것이 아닙니다. 특히 환경 설정, 정적 웹 자원, 로그 파일 등은 사용자에 의해 자주 읽고 쓰는 용도로 쓰입니다. 이 Docker 컨테이너에 있는 동적인 경로들만 따로 격리해서 관리하고 싶습니다. 그렇다면, Data Volume 을 이용하면 됩니다. /data/uhee-proxy/logs 경로로 디렉토리를 만든 후, 다음의 명령을 실행합니다.

1
2
3
$ docker stop uhee-proxy
$ docker run --name uhee-proxy --volume /data/uhee-proxy/logs:/var/log/nginx:rw --detach nginx
$ curl http://localhost:80
curl 로 access log 가 찍혔을 테니, /data/uhee-proxy/logs 로 이동합니다. access.log 가 생긴 것을 확인할 수 있습니다. /data/uhee-proxy/logs 디렉토리가 Data Volume 으로 uhee-proxy 컨테이너에 마운트된 것입니다. 이로써, 이 Docker 컨테이너에서는 /var/log/nginx 디렉토리를 호스트 OS 의 /data/uhee-proxy/logs 로 읽고 씁니다. 이러한 설정을 Nginx 환경 설정 파일과 정적 리소스들을 저장하는 디렉토리에도 함께 설정해주면 됩니다. 덧붙여, Data Volume 컨테이너라는 개념이 있는데, Docker 컨테이너를 따로 생성하고 이를 Docker 컨테이너 단위로 mount 해버릴 수 있습니다. 이 안에 환경 설정, 정적 웹자원, 로그 등을 격리 후 관리합니다. 이러한 비지속적인 데이터들을 Data Volume Container 단위로 묶어버리는 것도 고려할 만 합니다.