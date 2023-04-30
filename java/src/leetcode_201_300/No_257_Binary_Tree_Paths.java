package leetcode_201_300;

import utils.TreeNode;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2018/10/21
 */

public class No_257_Binary_Tree_Paths {
    private void deleteLast(StringBuilder sb, int len) {
        while (len-- != 0) {
            sb.deleteCharAt(sb.length() - 1);
        }
    }

    private void visit(TreeNode root, List<String> ans, StringBuilder sb) {
        String i = String.valueOf(root.val);
        sb.append(i);
        if (root.left == null && root.right == null) { ans.add(sb.toString()); }
        else {
            sb.append("->");
            if (root.left != null) { visit(root.left, ans, sb); }
            if (root.right != null) { visit(root.right, ans, sb); }
            deleteLast(sb, 2);
        }
        deleteLast(sb, i.length());
    }

    public List<String> binaryTreePaths(TreeNode root) {
        List<String> ans = new ArrayList<>();
        if (root == null) { return ans; }
        StringBuilder sb = new StringBuilder();
        visit(root, ans, sb);
        return ans;
    }
}
