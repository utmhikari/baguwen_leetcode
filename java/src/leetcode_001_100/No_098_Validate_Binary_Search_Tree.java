package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/15
 * 098. 判定二叉搜索树是否合法
 */
import java.util.List;
import java.util.ArrayList;
public class No_098_Validate_Binary_Search_Tree {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    public void genList(List<Integer> visitList, TreeNode root) {
        if(root.left != null) {
            genList(visitList, root.left);
        }
        visitList.add(root.val);
        if(root.right != null) {
            genList(visitList, root.right);
        }
    }

    public boolean isValidBST(TreeNode root) {
        if(root == null) {
            return true;
        }
        List<Integer> visitList = new ArrayList<>();
        genList(visitList, root);
        for(int i = 0; i < visitList.size() - 1; i++) {
            if(visitList.get(i) >= visitList.get(i + 1)) {
                return false;
            }
        }
        return true;
    }
}
