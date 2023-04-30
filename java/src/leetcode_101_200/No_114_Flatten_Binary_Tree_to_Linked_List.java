package leetcode_101_200;

import utils.TreeNode;
/**
 * Created by 36191 on 2017/12/10
 * 114.扁平化二叉树(右子树)
 */

public class No_114_Flatten_Binary_Tree_to_Linked_List {

    public void flatten(TreeNode root) {
        if(root != null) {
            flatten(root.right);
            flatten(root.left);
            if(root.left != null) {
                TreeNode temp = root.left;
                while(temp.right != null) {
                    temp = temp.right;
                }
                temp.right = root.right;
                root.right = root.left;
                root.left = null;
            }
        }
    }

    public static void main(String[] args) {
        No_114_Flatten_Binary_Tree_to_Linked_List solution = new No_114_Flatten_Binary_Tree_to_Linked_List();
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(3);
        root.right = new TreeNode(4);
        root.right.left = new TreeNode(5);
        solution.flatten(root);
        TreeNode.output(root);
    }
}
