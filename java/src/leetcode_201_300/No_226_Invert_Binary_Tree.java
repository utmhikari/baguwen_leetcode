package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/6
 */
import utils.*;
public class No_226_Invert_Binary_Tree {
    public TreeNode invertTree(TreeNode root) {
        if (root != null) {
            if (root.left != null) { root.left = invertTree(root.left); }
            if (root.right != null) { root.right = invertTree(root.right); }
            TreeNode tmp = root.left;
            root.left = root.right;
            root.right = tmp;
        }
        return root;
    }
}
