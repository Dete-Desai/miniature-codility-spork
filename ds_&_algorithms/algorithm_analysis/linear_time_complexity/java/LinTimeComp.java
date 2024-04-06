package algorithm_analysis.linear_time_complexity.java;

public class LinTimeComp {

    static int sumOperation(int n){
        int sum = 0;
        for (int i = 1; i<=n; i++){
            sum = sum + i; // 
        }
        return sum;
    }

    public static void main(String[] args){
        int val = 3;
        int finalSum = sumOperation(val);

        System.out.println("The value of the sum is:" +finalSum);
    }
    
}
