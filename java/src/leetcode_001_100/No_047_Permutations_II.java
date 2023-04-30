package leetcode_001_100;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Arrays;

/**
 * Created by 36191 on 2017/10/14
 * 047. 排列组合(包含重复数字)
 */

public class No_047_Permutations_II {

    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> list = new ArrayList<>();
        Arrays.sort(nums);
        boolean[] used = new boolean[nums.length];
        backtrack(list, new LinkedList<>(), nums, used);
        return list;
    }

    private void backtrack(List<List<Integer>> list, LinkedList<Integer> tempList, int[] nums, boolean[] used){
        if(tempList.size() == nums.length){
            ArrayList<Integer> ans = new ArrayList<>();
            for(int i : tempList) {
                ans.add(nums[i]);
            }
            list.add(ans);
        } else{
            for(int i = 0; i < nums.length; i++){
                if(used[i] || (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])) {
                    continue;
                }
                tempList.add(i);
                used[i] = true;
                backtrack(list, tempList, nums, used);
                used[i] = false;
                tempList.removeLast();
            }
        }
    }

    public void output(int[] nums) {
        List<List<Integer>> ans = permuteUnique(nums);
        for(List<Integer> a : ans) {
            for(int i : a) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_047_Permutations_II solution = new No_047_Permutations_II();
        solution.output(new int[] {1, 1, 3, 4});
    }


}
