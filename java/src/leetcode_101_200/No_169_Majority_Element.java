package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/25
 */
import java.util.*;

public class No_169_Majority_Element {
    public int majorityElement(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>(32);
        int th = nums.length / 2;
        for (int i : nums) {
            if (!map.containsKey(i)) { map.put(i, 1); }
            else { map.put(i, map.get(i) + 1); }
            if (map.get(i) > th) { return i; }
        }
        return -1;
    }

    public static void main(String[] args) {
        No_169_Majority_Element solution = new No_169_Majority_Element();
        int[][] inputs = new int[][] {
                new int[] {3, 2, 3},
                new int[] {2, 2, 2, 1, 1, 1, 2},
        };
        for (int[] input : inputs) {
            System.out.println(solution.majorityElement(input));
        }
    }
}
