package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/21
 */
import java.util.*;
import utils.*;

public class No_145_Binary_Tree_Postorder_Traversal {
    private void visit(List<Integer> post, TreeNode node) {
        if (node.left != null) { visit(post, node.left); }
        if (node.right != null) { visit(post, node.right); }
        post.add(node.val);
    }

    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> post = new ArrayList<>();
        if (root == null) { return post; }
        visit(post, root);
        return post;
    }
}
