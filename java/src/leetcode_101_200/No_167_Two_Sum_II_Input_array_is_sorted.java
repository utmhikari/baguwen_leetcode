package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */

public class No_167_Two_Sum_II_Input_array_is_sorted {

    public int[] twoSum(int[] numbers, int target) {
        for (int i = numbers.length - 1; i > 0; i--) {
            for (int j = 0; j < i; j++) {
                int sum = numbers[i] + numbers[j];
                if (sum == target) { return new int[] {j + 1, i + 1}; }
                else if (sum > target) { break; }
            }
        }
        return null;
    }
}
