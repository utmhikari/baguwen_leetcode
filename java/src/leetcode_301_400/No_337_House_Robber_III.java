package leetcode_301_400;

import utils.*;
/**
 * Created by 36191 on 2019/1/13
 */
public class No_337_House_Robber_III {
    
    public int rob(TreeNode root) {
        if (root == null) { return 0; }
        int left = root.left != null ? rob(root.left.left) + rob(root.left.right) : 0;
        int right = root.right != null ? rob(root.right.left) + rob(root.right.right) : 0;
        return Math.max(root.val + left + right, rob(root.left) + rob(root.right));
    }
}
