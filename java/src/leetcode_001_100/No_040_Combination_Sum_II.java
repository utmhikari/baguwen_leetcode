package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/12
 * 040. 和的组合II
 */
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Stack;

public class No_040_Combination_Sum_II {

    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> answers = new ArrayList<>();
        if(candidates.length == 0) {
            return answers;
        }
        Stack<Integer> answer = new Stack<>();
        int sum = 0;
        for(int j = 0;;) {
            if(!answer.isEmpty() && j < answer.peek()) {
                j = answer.peek();
            }
            if(j >= candidates.length) {
                if(answer.isEmpty()) {
                    break;
                } else {
                    sum -= candidates[answer.peek()];
                    j = answer.pop() + 1;
                    while(j < candidates.length && candidates[j] == candidates[j - 1]) {
                        j++;
                    }
                    continue;
                }
            }
            answer.push(j);
            sum += candidates[j++];

            if(sum >= target) {
                if(sum == target) {
                    List<Integer> ans = new ArrayList<>();
                    for(int a : answer) {
                        ans.add(candidates[a]);
                    }
                    answers.add(ans);
                }
                if(answer.size() == 1) {
                    break;
                } else {
                    sum -= candidates[answer.pop()];
                    sum -= candidates[answer.peek()];
                    j = answer.pop() + 1;
                    while(j < candidates.length && candidates[j] == candidates[j - 1]) {
                        j++;
                    }
                }
            }
        }
        return answers;
    }

    public void test(int[] candidates, int target) {
        List<List<Integer>> ans = combinationSum2(candidates, target);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_040_Combination_Sum_II solution = new No_040_Combination_Sum_II();
        solution.test(new int[] {10, 1, 2, 7, 6, 1, 5}, 8);
        solution.test(new int[] {2, 5, 2, 1, 2}, 5);
    }
}
