package leetcode_401_500;

import utils.TreeNode;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_404_Sum_of_Left_Leaves {
    private static int sum;

    private void traverse(TreeNode root, int dir) {
        if (root.left == null && root.right == null) {
            if (dir == -1) { sum += root.val; }
        } else {
            if (root.left != null) { traverse(root.left, -1); }
            if (root.right != null) { traverse(root.right, 1); }
        }
    }

    public int sumOfLeftLeaves(TreeNode root) {
        if (root == null) { return 0; }
        sum = 0;
        traverse(root, 0);
        return sum;
    }
}
