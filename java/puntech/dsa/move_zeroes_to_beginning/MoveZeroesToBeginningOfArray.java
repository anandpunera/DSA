
public class MoveZeroesToBeginningOfArray {

    public static void main(String[] args) {
        int[] input = new int[]{1, 0, 3, 4, 0, 6, 7, 0, 19};
        printNumbersZeroPrefixed(input);
    }

    private static void printNumbersZeroPrefixed(int[] input) {
        int lastZeroIndex = -1;

        for (int i = input.length - 1; i >= 0; i--) {
            if (input[i] == 0) {
                lastZeroIndex = i;
                break;
            }
        }

        for (int i = input.length - 2; i >= 0; i--) {
            if (input[i] > 0 && lastZeroIndex > 0) {
                input[lastZeroIndex] = input[i];
                lastZeroIndex--;
                input[i] = 0;
            }
        }

        for (int i = input.length - 1; i >= 0; i--) {
            System.out.println(input[i]);
        }
    }
}