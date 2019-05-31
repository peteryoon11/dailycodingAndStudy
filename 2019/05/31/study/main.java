package test.post;

import java.util.Arrays;
import java.util.List;

public class TestForTest {
	public static void main(String[] args) {
		Integer[] tempArray= {8,4,6,10};
		Integer[] tempArray2= {8,4,6};
		int count=4;
		
		//minimumTime(count,Arrays.asList(tempArray));
		System.out.println("result =  "+minimumTime(count,Arrays.asList(tempArray)));
		
		System.out.println("factorial "+factorial(4)); 
		makeArray(3,tempArray2);
	}
	/*
	 * 2가지
	 * 각 배열을 더할때 모든 값을 더하는 와중에 
	 * */
	static int minimumTime(int count, List<Integer> array) {
		int[] flow = new int[array.size()-1];
		int total =0;
		int temp=0;
		array.toArray();
		for(Integer item : array) {
			total+=item;
			if(array.get(0)!=item) {
				temp+=total;	
			}
			
			System.out.println(" now temp = "+ temp+" total = "+total);
			
		}
		System.out.println("temp = "+ temp);
		Arrays.sort(array.toArray()	);
		
		return total;
	}
	static List<List<Integer>> makeArray(int num , Integer[] tempArray){
		/*
		 * */
		
		int[][] flow = new int[factorial(num)][num];
		//return Arrays.asList(flow);
		int count =0;
		int subcount=tempArray.length-1;
		for(int i=0;flow.length>i;i++) { //
			
			
			for(int j=0;flow[0].length>j;j++) {
			/*	
				if(subcount>0) {
					System.out.print(" subcount ");
					subcount--;
				}else {
					System.out.print(" subcount ");
				}
				*/
				
				if(i==j){
					System.out.print("A "+i+" ");
					flow[i][flow[i][j]]=tempArray[i];	
				}else {
					System.out.print("B "+j+" ");
					flow[i][flow[i][j]]=tempArray[j];	
				}
				
				count++;
			}
			System.out.println();
		}
		// 출력 용 
		for(int i=0;flow.length>i;i++) { //
			System.out.print("{ ");
			for(int j=0;flow[0].length>j;j++) {
				
				System.out.print(flow[i][j]+"( "+i+", "+j+" ),");
			}
			System.out.println("} ");
		}
		
		System.out.println("count = " + count + " total count = "+flow.length+ " sub count "+flow[0].length);
		return null;
	}
	
	static int factorial(int num) {
		int result=1;
		for(int i=num;i>1;i--) {
			result*=i;
		}
		
		return result;
	}
}
