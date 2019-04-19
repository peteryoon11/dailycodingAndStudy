#  select in  limit 1000 object 


select in 을 쓰는데 가끔(정확히는 그런 가능성에 대한 예외 처리..) 찾아야 하는 경우가 있어서 코드를 추가 했다.

다른 좋은 방법이 보이거나 생간 날때까지는 그냥 그대로 사용 할듯 하다.



				List<String> names = Arrays.asList("네이버", "블로그", "부르곰", "블로그", "박성균");

				  
 int loopCount = names.size()/2;
				  
				  for(int i =0;i<=loopCount;i++) {
					  
					if(i==loopCount) {
						System.out.println(names.subList(2*(i), names.size()));
					}else {
						System.out.println(names.subList(2*(i), 2*(i+1)));
					}
					  
					  
				  } 