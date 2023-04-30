package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/28
 * No.193 二叉树右视图
 */

import java.util.*;
import utils.*;

public class No_193_Binary_Tree_Right_Side_View {

    private void traverse(List<TreeNode> queue, List<Integer> rsv) {
        if (queue.size() != 0) {
            rsv.add(queue.get(queue.size() - 1).val);
            List<TreeNode> tmp = new ArrayList<>();
            for (TreeNode tn : queue) {
                if (tn.left != null) { tmp.add(tn.left); }
                if (tn.right != null) { tmp.add(tn.right); }
            }
            queue = tmp;
            traverse(queue, rsv);
        }
    }

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> rsv = new ArrayList<>();
        if (root == null) { return rsv; }
        List<TreeNode> queue = new ArrayList<TreeNode>() {{ add(root); }};
        traverse(queue, rsv);
        return rsv;
    }
}
