package arrays.java;

public class adding_and_updating_arrays{
    public void createArray(int[] array){
        int n = array.length;

        for(int i = 0; i<n; i++){
            System.out.println("The value at index" + i + "of the array is: " + array[i] + "");
        }
        System.out.println();
    }

    public void addAndUpdateArray(){
        int[] myArray = new int[5];
        // the first elements in the array before adding
        //createArray(myArray);
        // adding elements into the array
        myArray[0] = 1;
        myArray[1] = 1;
        myArray[2] = 9;
        myArray[3] = 9;
        myArray[4] = 5;
        //createArray(myArray);
        // updating elements into the array
        myArray[4] = 1;
        myArray[3] = 1;
        myArray[2] = 9;
        myArray[1] = 9;
        myArray[0] = 5;
        createArray(myArray);
    }
    public static void main(String[] args){
        adding_and_updating_arrays addingAndUpdatingArrays = new adding_and_updating_arrays();
        addingAndUpdatingArrays.addAndUpdateArray();
    }
}

