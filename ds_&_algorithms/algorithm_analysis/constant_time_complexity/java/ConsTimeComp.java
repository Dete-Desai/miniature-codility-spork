package algorithm_analysis.constant_time_complexity.java;

public class ConsTimeComp {

    static int sum(int x, int y){
        int result = x+y; // Constant time operation
        return result;
    }

    public static void main(String[] args){
        int adder = 10;
        int addee = 10;
        int finalsSum = sum(addee, adder);
        System.out.println("The value of sum is:" +finalsSum);
    }
    
}
