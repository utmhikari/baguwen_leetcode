package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/13
 */

import utils.*;

public class No_222_Count_Complete_Tree_Nodes {

    public int countNodes(TreeNode root) {
        if (root == null) { return 0; }
        TreeNode left = root.left, right = root.right;
        int sum = 2;
        while (left != null && right != null) {
            left = left.left; right = right.right; sum *= 2;
        }
        return left == right ? sum - 1 : 1 + countNodes(root.left) + countNodes(root.right);
    }
}
