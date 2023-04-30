package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/4
 * 106.按中序遍历与后序遍历生成二叉树
 */
import utils.TreeNode;

public class No_106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal {

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        return helper(0, postorder.length - 1, 0, inorder.length - 1, postorder, inorder);
    }

    public TreeNode helper(int postStart, int postEnd, int inStart, int inEnd, int[] postorder, int[] inorder) {
        if (postStart > postorder.length - 1 || inStart > inEnd) {
            return null;
        }
        TreeNode root = new TreeNode(postorder[postEnd]);
        int inIndex = 0;
        for (int i = inStart; i <= inEnd; i++) {
            if (inorder[i] == root.val) {
                inIndex = i;
                break;
            }
        }
        System.out.println(postStart + ", " + postEnd + ", " + inStart + ", " + inEnd + ", " + inIndex);
        root.left = helper(postStart, postStart + inIndex - inStart - 1, inStart, inIndex - 1, postorder, inorder);
        root.right = helper(inIndex, postEnd - 1, inIndex + 1, inEnd, postorder, inorder);
        return root;

    }

    public static void main(String[] args) {
        No_106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal solution = new No_106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal();
        solution.buildTree(new int[] {1, 3, 2}, new int[] {3, 2, 1});
    }
}
