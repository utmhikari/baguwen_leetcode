package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/12
 * 039. 和的组合
 */
import java.util.List;
import java.util.Stack;
import java.util.ArrayList;
import java.util.Arrays;

public class No_039_Combination_Sum {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
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
                    continue;
                }
            }
            answer.push(j);
            sum += candidates[j];
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
                }
            }
        }
        return answers;
    }

    public void test(int[] candidates, int target) {
        List<List<Integer>> ans = combinationSum(candidates, target);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
        System.out.println();
    }
    public static void main(String[] args) {
        No_039_Combination_Sum solution = new No_039_Combination_Sum();
        solution.test(new int[] {2, 3, 6, 7}, 7);
        solution.test(new int[] {1, 2}, 4);
    }
}
