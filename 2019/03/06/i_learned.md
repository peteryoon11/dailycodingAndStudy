## java 8 의 람다 기능과 com.google.collections 를 이용한 자료구조 정리 



List<DTOTemp> listemp=new ArrayList<>();
	///	listemp.add(new DTOTemp("1", "1"));
		listemp.add(new DTOTemp("11", "2"));
		listemp.add(new DTOTemp("11", "3"));
		listemp.add(new DTOTemp("11", "4"));
		
		listemp.add(new DTOTemp("22", "5"));
		listemp.add(new DTOTemp("22", "6"));
		listemp.add(new DTOTemp("33", "7"));

        String[] array1 = {"name1","name1","name2","name2", "name2"};
		 List<String> asList = Arrays.asList(array1);
		 Set<String> mySet = new HashSet<String>(asList);
		 for(String s: mySet){

		  System.out.println(s + " " +Collections.frequency(asList,s));

		 }
		 //Maps.filterValues(unfiltered, valuePredicate)
		 for(DTOTemp item : listemp) {
			 System.out.println(item.getAb() +" "+item.getBb());
		 }