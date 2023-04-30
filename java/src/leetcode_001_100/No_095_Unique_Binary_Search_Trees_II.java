package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/11
 * 095. 生成二叉搜索树2
 */
import java.util.List;
import java.util.ArrayList;

public class No_095_Unique_Binary_Search_Trees_II {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    /*
    https://discuss.leetcode.com/topic/3079/a-simple-recursive-solution
     */

    public List<TreeNode> generateTrees(int n) {
        if(n == 0) {
            return new ArrayList<TreeNode>();
        }
        return genTrees(1, n);
    }

    public List<TreeNode> genTrees(int start, int end) {
        List<TreeNode> list = new ArrayList<>();
        if(start > end) {
            list.add(null);
            return list;
        }
        if(start == end) {
            list.add(new TreeNode(start));
            return list;
        }
        List<TreeNode> left, right;
        for(int i = start; i <= end; i++) {
            left = genTrees(start, i - 1);
            right = genTrees(i + 1,end);
            for(TreeNode lnode: left) {
                for(TreeNode rnode: right) {
                    TreeNode root = new TreeNode(i);
                    root.left = lnode;
                    root.right = rnode;
                    list.add(root);
                }
            }
        }
        return list;
    }

    public void output(TreeNode root) {
        if(root != null) {
            System.out.print(root.val + ", ");
            output(root.left);
            output(root.right);
        }
    }

    public void test(int n) {
        List<TreeNode> ans = generateTrees(n);
        for(TreeNode root : ans) {
            output(root);
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_095_Unique_Binary_Search_Trees_II solution = new No_095_Unique_Binary_Search_Trees_II();
        solution.test(4);
    }
}
