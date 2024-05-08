package divide_and_conquer_algorithm.java;

public class divide_and_conquer_algorithm {
    static int recursiveFuction(int n){
        if (n != 0){
            return n * recursiveFuction(n- 1);
        } else {
            return 1;
        }
    } 

    public static void main(String[] args){
        int number = 4;
        int result = recursiveFuction(number);
        System.out.println(number + " & factorial = " + result);
    }
}
