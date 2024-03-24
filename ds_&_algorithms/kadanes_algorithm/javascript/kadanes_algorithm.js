const maxSubarraySum=(array)=>{
    let max_ending_here = 0;
    let max_so_far = 0;

    for (let i=0; i<array.length; i++){
        max_ending_here += array[i];

        if(max_ending_here<0){
            max_ending_here=0
        }
        if(max_so_far<max_ending_here){
            max_so_far = max_ending_here
        }
    }
    return max_so_far
}

const input_array = [-2, 1, -3, 4, -1, 2, 1, -5, 4];

const max_sum = maxSubarraySum(input_array);

console.log(`The maximum sum of the sub arrays in the give array is ${max_sum}`);