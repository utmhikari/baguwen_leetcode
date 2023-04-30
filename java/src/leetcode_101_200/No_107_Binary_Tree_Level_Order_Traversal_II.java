package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/5
 * 107.从底到上遍历二叉树
 */
import java.util.List;
import java.util.ArrayList;
import java.util.LinkedList;
public class No_107_Binary_Tree_Level_Order_Traversal_II {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    private void traverse(TreeNode node, int level, LinkedList<List<Integer>> ans) {
        if(node != null) {
            if(ans.size() <= level) {
                ans.add(new ArrayList<>());
            }
            ans.get(level).add(node.val);
            traverse(node.left, level + 1, ans);
            traverse(node.right, level + 1, ans);
        }
    }

    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        LinkedList<List<Integer>> ans = new LinkedList<>();
        traverse(root, 0, ans);
        ArrayList<List<Integer>> newAns = new ArrayList<>();
        while(!ans.isEmpty()) {
            newAns.add(ans.getLast());
            ans.removeLast();
        }
        return newAns;
    }
}
