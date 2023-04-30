package leetcode_101_200;

import utils.TreeNode;
import java.util.LinkedList;

/**
 * Created by 36191 on 2017/12/9
 * 111.二叉树最小深度
 */

public class No_111_Minimum_Depth_of_Binary_Tree {

    private class NodeDepth {
        TreeNode node;
        int depth;
        NodeDepth(TreeNode node, int depth) {
            this.node = node;
            this.depth = depth;
        }
    }

    private static int minDepth = 0;

    private void traverse(LinkedList<NodeDepth> visitList) {
        if (visitList.size() > 0) {
            if (visitList.getFirst().node.left == null && visitList.getFirst().node.right == null) {
                minDepth = visitList.getFirst().depth;

            } else {
                if (visitList.getFirst().node.left != null) {
                    visitList.add(new NodeDepth(visitList.getFirst().node.left, visitList.getFirst().depth + 1));
                }
                if (visitList.getFirst().node.right != null) {
                    visitList.add(new NodeDepth(visitList.getFirst().node.right, visitList.getFirst().depth + 1));
                }
                visitList.removeFirst();
                if(!visitList.isEmpty()) {
                    traverse(visitList);
                }
            }
        }
    }

    public int minDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        LinkedList<NodeDepth> visitList = new LinkedList<>();
        visitList.add(new NodeDepth(root, 1));
        minDepth = 0;
        traverse(visitList);
        return minDepth;
    }

    public static void main(String[] args) {
        No_111_Minimum_Depth_of_Binary_Tree solution = new No_111_Minimum_Depth_of_Binary_Tree();
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        System.out.println(solution.minDepth(root));
    }
}
