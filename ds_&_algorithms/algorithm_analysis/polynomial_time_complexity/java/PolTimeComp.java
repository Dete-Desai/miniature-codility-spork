package algorithm_analysis.polynomial_time_complexity.java;

public class PolTimeComp {
    public void printFunc(int n) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                System.out.println("i = " + i + ", j = " + j);
            }
            System.out.println("End of inner loop");
        }
        System.out.println("End of outer loop");
    }

    public static void main(String[] args) {
        int value = 3;
        PolTimeComp obj = new PolTimeComp(); // Creating an instance of the class
        System.out.println("The value of the range is as follows:");
        obj.printFunc(value); // Correct method call
    }
}
