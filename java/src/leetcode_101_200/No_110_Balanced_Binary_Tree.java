package leetcode_101_200;

import utils.TreeNode;
import java.util.HashSet;
/**
 * Created by 36191 on 2017/12/8
 * 110.二叉树是否平衡
 */

public class No_110_Balanced_Binary_Tree {

    public boolean isBalanced(TreeNode root) {
        if(root == null){
            return true;
        }
        return height(root) != -1;
    }

    public int height(TreeNode node){
        if(node == null){
            return 0;
        }
        int lH = height(node.left);
        if(lH == -1){
            return -1;
        }
        int rH = height(node.right);
        if(rH == -1){
            return -1;
        }
        if(lH - rH < -1 || lH - rH > 1){
            return -1;
        }
        return Math.max(lH, rH) + 1;
    }
}
