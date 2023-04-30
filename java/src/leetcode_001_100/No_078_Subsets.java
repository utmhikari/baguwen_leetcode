package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/21
 * 078.寻找所有子集
 */
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
public class No_078_Subsets {

    public void backtrack(List<List<Integer>> subsets, List<Integer> subset, int[] nums, int start) {
        for(int i = start; i < nums.length; i++) {
            subset.add(nums[i]);
            subsets.add(new ArrayList<>(subset));
            backtrack(subsets, subset, nums, i + 1);
            subset.remove(subset.size() - 1);
        }
    }

    public List<List<Integer>> subsets(int[] nums) {
        LinkedList<List<Integer>> subsets = new LinkedList<>();
        subsets.add(new ArrayList<>());
        if(nums.length == 0) {
            return subsets;
        }
        ArrayList<Integer> subset = new ArrayList<>();
        backtrack(subsets, subset, nums, 0);
        return subsets;
    }

    public void test(int[] nums) {
        List<List<Integer>> ans = subsets(nums);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(i);
                System.out.print(", ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_078_Subsets solution = new No_078_Subsets();
        solution.test(new int[] {1, 2, 3, 4});
    }
}
