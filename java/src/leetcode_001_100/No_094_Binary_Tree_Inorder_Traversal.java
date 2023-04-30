package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/8
 * 094. 中序遍历
 */
import java.util.List;
import java.util.ArrayList;
public class No_094_Binary_Tree_Inorder_Traversal {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    public void traversal(List<Integer> visit, TreeNode root) {
        if(root == null) {
            return;
        }
        traversal(visit, root.left);
        visit.add(root.val);
        traversal(visit, root.right);
    }

    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> visit = new ArrayList<>();
        traversal(visit, root);
        return visit;
    }

    public void test() {
        TreeNode t1 = new TreeNode(1);
        t1.right = new TreeNode(2);
        t1.right.left = new TreeNode(3);
        List<Integer> ans1 = inorderTraversal(t1);
        for(int i : ans1) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }
    public static void main(String[] args) {
        No_094_Binary_Tree_Inorder_Traversal solution = new No_094_Binary_Tree_Inorder_Traversal();
        solution.test();
    }
}
