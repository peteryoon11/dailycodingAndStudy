git add . && git commit -m 'add content for today' && git push origin master

CORS 에 대한 정리 


https://www.zerocho.com/category/NodeJS/post/5a6c347382ee09001b91fb6a



웹 브라우저에서 콘솔 창에서 쳐보면 알수있음 

var xhr = new XMLHttpRequest();
xhr.onload = function() {
   console.log('xhr loaded');
};
xhr.open('GET', 'https://stackoverflow.com/');
xhr.send();