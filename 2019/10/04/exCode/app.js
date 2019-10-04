const shiftArray=(array, count)=>{

    let result=[];
    // array[0:count];
    const preArray=array.slice(0,count);
    const postArray=array.slice(count,array.length);
   
    for(const item of postArray){
        result.push(item);
    }
    for(const item of preArray){
        result.push(item);
    }

    // result=[preArray,postArray];
    return result;

};

const count=2;
const array=[1,2,3,4,5];
console.log(array.slice(count,array.length));
console.log(array.slice(0,count));


console.log(array.slice(count,array.length));
console.log(array.slice(count,array.length));




console.log(shiftArray(array,count));