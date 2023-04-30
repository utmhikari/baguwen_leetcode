package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/15
 */
import utils.*;
public class No_235_Lowest_Common_Ancestor_of_a_Binary_Search_Tree {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) { return null; }
        if (p.val == root.val || q.val == root.val) { return root; }
        if ((p.val < root.val && root.val < q.val) || (p.val > root.val && root.val > q.val)) {
            return root;
        } else {
            if (p.val > root.val) { return lowestCommonAncestor(root.right, p, q); }
            return lowestCommonAncestor(root.left, p, q);
        }
    }
}
