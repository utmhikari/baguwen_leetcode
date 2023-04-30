package leetcode_401_500;

import utils.TreeNode;
/**
 * Created by 36191 on 2019/2/12
 */
public class No_437_Path_Sum_III {

    private int pathSumChild(TreeNode child, int sum) {
        if (child == null) { return 0; }
        int newSum = sum - child.val;
        int ans = newSum == 0 ? 1 : 0;
        return ans + pathSumChild(child.left, newSum) + pathSumChild(child.right, newSum);
    }

    public int pathSum(TreeNode root, int sum) {
        if (root == null) { return 0; }
        return pathSum(root.left, sum) + pathSum(root.right, sum) + pathSumChild(root, sum);
    }

}
