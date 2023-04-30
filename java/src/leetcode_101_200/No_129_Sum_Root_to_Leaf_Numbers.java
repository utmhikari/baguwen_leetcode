package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/9
 */

import utils.TreeNode;
public class No_129_Sum_Root_to_Leaf_Numbers {
    private static int SUM = 0;

    private void sumNumbers(TreeNode root, int tmp) {
        if (root.left == null && root.right == null) {
            SUM += tmp;
        } else {
            tmp *= 10;
            if (root.left != null) { sumNumbers(root.left, tmp + root.left.val); }
            if (root.right != null) { sumNumbers(root.right, tmp + root.right.val); }
        }
    }

    public int sumNumbers(TreeNode root) {
        if(root == null) { return 0; }
        SUM = 0;
        sumNumbers(root, root.val);
        return SUM;
    }
}
