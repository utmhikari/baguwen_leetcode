package leetcode_101_200;

import java.util.List;
import utils.TreeNode;
import java.util.ArrayList;
/**
 * Created by 36191 on 2017/12/10
 * 113.加到底等于特定数的答案列表
 */

public class No_113_Path_Sum_II {

    public void pathSum(TreeNode node, int sum, int tempSum, ArrayList<List<Integer>> answers, ArrayList<Integer> ans) {
        ans.add(node.val);
        tempSum += node.val;
        if(node.left == null && node.right == null) {
            if(tempSum == sum) {
                answers.add(new ArrayList<>(ans));
            }
        } else {
            if(node.left != null) {
                pathSum(node.left, sum, tempSum, answers, ans);
            }
            if(node.right != null) {
                pathSum(node.right, sum, tempSum, answers, ans);
            }
        }
        ans.remove(ans.size() - 1);
    }

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        ArrayList<List<Integer>> answers = new ArrayList<>();
        ArrayList<Integer> ans = new ArrayList<>();
        if(root == null) {
            return answers;
        }
        int tempSum = 0;
        pathSum(root, sum, tempSum, answers, ans);
        return answers;
    }
}
