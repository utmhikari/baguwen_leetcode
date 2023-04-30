package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/21
 */
import java.util.*;
import utils.*;

public class No_144_Binary_Tree_Preorder_Traversal {
    private void visit(List<Integer> po, TreeNode root) {
        po.add(root.val);
        if (root.left != null) { visit(po, root.left); }
        if (root.right != null) { visit(po, root.right); }
    }

    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> po = new ArrayList<>();
        if (root == null) { return po; }
        visit(po, root);
        return po;
    }
}
