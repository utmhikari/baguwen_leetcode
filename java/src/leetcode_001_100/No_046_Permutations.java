package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/14
 * 046. 所有排列组合
 */
import java.util.List;
import java.util.ArrayList;
import java.util.LinkedList;

public class No_046_Permutations {

    /* use backtrack!!! */

    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> list = new ArrayList<>();
        backtrack(list, new LinkedList<>(), nums);
        return list;
    }

    private void backtrack(List<List<Integer>> list, LinkedList<Integer> tempList, int [] nums){
        if(tempList.size() == nums.length){
            list.add(new LinkedList<>(tempList));
        } else{
            for(int n : nums){
                if(tempList.contains(n)) {
                    continue; // element already exists, skip
                }
                tempList.add(n);
                backtrack(list, tempList, nums);
                tempList.removeLast();
            }
        }
    }

    /*
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> answers = new ArrayList<>();
        int len = nums.length;
        if(len == 0) {
            return answers;
        }
        if(len == 1) {
            ArrayList<Integer> ans = new ArrayList<Integer>() {{ add(nums[0]); }};
            answers.add(ans);
            return answers;
        }
        Stack<Integer> answer = new Stack<>();
        LinkedList<Integer> tempList = new LinkedList<>();
        int tempIndex = 0;
        int count = 0;

        for(int i = 0; i < len; i++) {
            answer.push(i);
            count++;
        }

        while(true) {
            if(count == len) {
                List<Integer> init = new ArrayList<>();
                for(int i : answer) {
                    init.add(nums[i]);
                    System.out.print(nums[i] + ", ");
                }
                System.out.println();
                answers.add(init);
                tempList.add(answer.pop());
                count--;
                tempIndex = answer.peek() + 1;
                while(count != 0 && tempIndex >= len - count) {
                    System.out.println(answer.peek());
                    tempList.add(answer.pop());
                    count--;
                }
                if(count != 0) {
                    answer.push(tempList.getFirst());
                    tempList.removeFirst();
                    count++;
                }
            } else if(count == 0) {
                if(tempIndex == len) {
                    break;
                } else {
                    answer.push(tempList.getFirst());
                    tempList.removeFirst();
                    count++;
                }
            } else {
                answer.push(tempList.getLast());
                tempList.removeLast();
                count++;
            }
        }
        return answers;
    }
    */

    public void output(int[] nums) {
        List<List<Integer>> ans = permute(nums);
        for(List<Integer> a : ans) {
            for(int i : a) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
    }
    public static void main(String[] args) {
        No_046_Permutations solution = new No_046_Permutations();
        solution.output(new int[] {0, 1, 2});
        solution.output(new int[] {1, 2, 3, 4});
    }
}
