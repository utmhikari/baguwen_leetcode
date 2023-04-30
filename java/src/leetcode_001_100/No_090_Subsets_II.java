package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/31
 * 090. 子集（包含重复元素）
 */
import java.util.List;
import java.util.ArrayList;
import java.util.HashSet;
public class No_090_Subsets_II {
    private static int hash = 0;

    private void dfs(List<List<Integer>> answers, ArrayList<Integer> ans, HashSet<Integer> hashes, int hash,
                    int[] nums, int len, int start) {
        for(int i = start; i < len; i++) {
            ans.add(nums[i]);
            int hashAdd = nums[i] * (nums[i] + 7) * (nums[i] + 31) * (nums[i] + 127) + 1;
            hash += hashAdd;
            if(!hashes.contains(hash)) {
                answers.add(new ArrayList<>(ans));
                hashes.add(hash);
                dfs(answers, ans, hashes, hash, nums, len, i + 1);
            }
            ans.remove(ans.size() - 1);
            hash -= hashAdd;
        }
    }

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        ArrayList<List<Integer>> answers = new ArrayList<>();
        HashSet<Integer> hashes = new HashSet<>();
        answers.add(new ArrayList<Integer>());
        hashes.add(0);
        int len = nums.length;
        ArrayList<Integer> ans = new ArrayList<>();
        dfs(answers, ans, hashes, hash, nums, len, 0);
        hash = 0;
        return answers;
    }

    public void test(int[] nums) {
        List<List<Integer>> ans = subsetsWithDup(nums);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(i);
                System.out.print(", ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_090_Subsets_II solution = new No_090_Subsets_II();
        solution.test(new int[] {1, 2, 2});
    }
}
