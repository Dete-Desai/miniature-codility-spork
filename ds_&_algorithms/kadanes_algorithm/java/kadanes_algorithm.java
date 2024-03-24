package kadanes_algorithm.java;

public class kadanes_algorithm {
    static  int maxSubarraySum(int[] array){
        int max_ending_here = 0;
        int max_so_far = 0;

        for (int i=0; i< array.length; i++){
            max_ending_here = max_ending_here + array[i];

            if(max_ending_here < 0){
                max_ending_here = 0;
            }
            if(max_so_far < max_ending_here){
                max_so_far = max_ending_here;
            }
        }
        return max_so_far;
    }

    public static void main(String[] args){
        int[] input_array = {-2, 1, -3, 4, -1, 2, 1, -5, 4}; 
        int max_sum = maxSubarraySum(input_array);

        System.out.println("The value of the maximum sum of sub arrays in in the given array is:" + max_sum);
    }
}
