/**/


let answer="YNYYNN";
const main=()=>{
    let aArray=[1,2,3,4,5,6,7,8];
    let bArray=[1,2,3,4,9,10,11,12];
    let cArray=[1,2,3,4,9,10,13,14];
    let dArray=[1,2,3,4,9,11,13,15];
    

for(let i=0;i<16;i++){
    if(check(aArray, i)!=answer.substring(0,1))
    {
        continue;
    }
    return i;
}
return 0;
    
}


const check= (x, number)=>{
    for(const item of x){
        if(item===number){
            return 'Y';
        }else{
            return 'N';
        }
        
    }
}
console.log(main());