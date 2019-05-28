ref 
http://think-devops.com/blogs/load-balancing-nginx-haproxy.html

HAProxy and nginx can be configured together to work as an SSL offloader and a load balancer. Listed below are the steps to achieve the same on a centOS instance.

Assume 192.168.1.1 and 192.168.1.2 running web servers on port 80 and 192.168.1.3 running haproxy on port 8181.

CONFIGURING HAPROXY
Install HAProxy.
yum install -y haproxy
Edit the haproxy configuration to update the backend web servers and keep it at the basic log level.

global
	log 127.0.0.1   local2
	chroot          /var/lib/haproxy
	pidfile         /var/run/haproxy.pid
	maxconn         4096
	user            haproxy
	group           haproxy
	daemon

defaults
	mode            http
	log             global
	option          httplog
	option          dontlognull
	option          http-server-close
	option          forwardfor except 127.0.0.0/8
	option          redispatch
	retries         3
	timeout         http-request    20s
	timeout         queue           1m
	timeout         connect         10s
	timeout         client          1m
	timeout         server          1m
	timeout         http-keep-alive 30s
	timeout         check           10s
	maxconn         3000

frontend fe_http
	option          forwardfor except 127.0.0.1
	option          httpclose
	bind            *:8181
	default_backend be_http

backend  be_http
	balance         roundrobin
	option          httpchk
	server          ws_1 192.168.1.1:80 check port 80
	server          ws_2 192.168.1.2:80 check port 80
HAProxy doesn't start logging on installation, it uses syslog for the same. To enable logging install rsyslog and add a config for HAProxy
yum install -y rsyslog
Create a config file for haproxy logging /etc/rsyslog.d/haproxy.conf
$ModLoad imudp
$UDPServerRun 514
$template Haproxy,"%msg%\n"
local2.info -/var/log/haproxy.log;Haproxy
local2.notice -/var/log/haproxy.admin;Haproxy
# don't log anywhere else
local2.* ~
Edit /etc/sysconfig/rsyslog as below
SYSLOGD_OPTIONS="-c 2 -r"
Restart rsyslog and haproxy services. That sets up haproxy to bind to 8181 and check the ports 80 for all the backend web servers that would be load balanced.
CONFIGURING NGINX
Install nginx.
yum install -y nginx
Assuming the web servers would bind to default port 80, remove the default.conf in /etc/nginx/conf.d/. Copy the SSL certificate and private key to /etc/nginx/. Make sure to change the owner to nginx:nginx with mode 600 and 644 respectively.
Configure the default nginx server as below
error_log   /var/log/nginx/ssl_error.log debug;
access_log  /var/log/nginx/ssl_access.log;

upstream haproxy {
	server 192.168.1.3:8181;
}

server {
	listen 443 ssl;
	ssl_certificate         server.crt
	ssl_certificate_key     server_cert.key
	server_name             domain.com
	location / {
		proxy_pass            http://haproxy/;
		proxy_set_header      X-NginX-Proxy true;
		proxy_set_header      Host $http_host;
		proxy_set_header      X-Real-IP  $remote_addr;
		proxy_set_header      X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header      X-Forwarded-Proto https;

		proxy_redirect        default;
		proxy_redirect        http://$host/ https://$host/;
		proxy_redirect        http://hostname/ https://$host/;
		proxy_read_timeout    15s;
		proxy_connect_timeout 15s;
	}
	location ~ /\. { deny  all; }
}

server {
	listen    80;
	return    301 https://$host$request_uri;
}
Restart the nginx and haproxy services.