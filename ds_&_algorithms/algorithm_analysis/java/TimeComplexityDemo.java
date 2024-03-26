package algorithm_analysis.java;

public class TimeComplexityDemo {

    public int findSumSimplify(int n){
        return n* (n+1)/2;
    }

    // public int findSum(int n){
    //     int sum = 0;
    //     for(int i = 0; i < n; i++){
    //         sum = sum + i;
    //     }

    //     return sum;
    // }

    public static void main(String[] args){
        double now = System.currentTimeMillis();

        TimeComplexityDemo demo = new TimeComplexityDemo();
        //System.out.println(demo.findSum(99999));
        System.out.println(demo.findSumSimplify(99999));

        System.out.println("Time taken is: " + (System.currentTimeMillis() - now) + " milliseconds.");
    }
}