const  recursiveFunction = (n) => {
    if(n != 0)
        return n * recursiveFunction(n-1);
    else
       return 1;
}

const  number = 4;
const result = recursiveFunction(number);

console.log(`${number} & factorial = ${result}`);