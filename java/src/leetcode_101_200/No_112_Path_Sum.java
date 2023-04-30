package leetcode_101_200;
import utils.TreeNode;
/**
 * Created by 36191 on 2017/12/10
 * 112.加到底等于特定数
 */

public class No_112_Path_Sum {

    public boolean hasPathSum(TreeNode node, int sum, int tempSum) {
        tempSum += node.val;
        if(node.left == null && node.right == null) {
            return tempSum == sum;
        } else {
            boolean left = false;
            boolean right = false;
            if(node.left != null) {
                left = hasPathSum(node.left, sum, tempSum);
            }
            if(node.right != null) {
                right = hasPathSum(node.right, sum, tempSum);
            }
            return left || right;
        }
    }

    public boolean hasPathSum(TreeNode root, int sum) {
        if(root == null) {
            return false;
        }
        int tempSum = 0;
        return hasPathSum(root, sum, tempSum);
    }
}
