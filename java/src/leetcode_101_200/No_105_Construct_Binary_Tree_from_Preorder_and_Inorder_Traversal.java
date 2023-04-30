package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/1
 * 105.按前序与中序遍历创建二叉树
 */

public class No_105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    /* https://discuss.leetcode.com/topic/3695/my-accepted-java-solution */

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return helper(0, 0, inorder.length - 1, preorder, inorder);
    }

    public TreeNode helper(int preStart, int inStart, int inEnd, int[] preorder, int[] inorder) {
        if (preStart > preorder.length - 1 || inStart > inEnd) {
            return null;
        }
        TreeNode root = new TreeNode(preorder[preStart]);
        int inIndex = 0;
        for (int i = inStart; i <= inEnd; i++) {
            if (inorder[i] == root.val) {
                inIndex = i;
            }
        }
        root.left = helper(preStart + 1, inStart, inIndex - 1, preorder, inorder);
        root.right = helper(preStart + inIndex - inStart + 1, inIndex + 1, inEnd, preorder, inorder);
        return root;
    }



    /*
    private static utils.TreeNode root;

    private static void buildTree(utils.TreeNode node, int[] preorder, int[] inorder, int pStart, int pEnd, int iStart, int iEnd) {
        System.out.println(pStart + ", " + pEnd + ", " + iStart + ", " + iEnd);
        node.val = preorder[pStart];
        int iIndex = pStart;
        while(iIndex <= iEnd) {
            if(inorder[iIndex] == preorder[pStart]) {
                break;
            }
            iIndex++;
        }
        int leftLen = iIndex - iStart;
        int rightLen = iEnd - iIndex;
        if(pStart + 1 <= pStart + leftLen - 1 && pStart + 1 < preorder.length) {
            System.out.println("left");
            node.left = new utils.TreeNode(Integer.MIN_VALUE);
            buildTree(node.left, preorder, inorder, pStart + 1, pStart + leftLen, iStart, iIndex - 1);
        }
        if(pEnd - rightLen + 1 <= pEnd && pEnd - rightLen + 1 > pStart && pEnd < preorder.length) {
            System.out.println("right");
            node.right = new utils.TreeNode(Integer.MIN_VALUE);
            buildTree(node.right, preorder, inorder, pEnd - rightLen + 1, pEnd, iIndex + 1, iEnd);
        }
    }

    public utils.TreeNode buildTree(int[] preorder, int[] inorder) {
        if(preorder.length == 0 || preorder.length != inorder.length) {
            return null;
        }
        root = new utils.TreeNode(Integer.MIN_VALUE);
        buildTree(root, preorder, inorder, 0, preorder.length - 1, 0, inorder.length - 1);
        return root;
    }
    */

    public static void main(String[] args) {
        No_105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal solution = new No_105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal();
        solution.buildTree(new int[] {6, 1, 2, 4, 3, 5}, new int[] {1, 2, 3, 4, 5, 6});
    }
}
