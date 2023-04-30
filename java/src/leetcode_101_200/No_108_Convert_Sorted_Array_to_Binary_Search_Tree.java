package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/5
 * 108.将排序后队列转换成二叉搜索树
 */
import utils.TreeNode;
public class No_108_Convert_Sorted_Array_to_Binary_Search_Tree {
    private TreeNode generateBST(int[] nums, int start, int end) {
        if(start > end) {
            return null;
        }
        int half = start + (end - start) / 2;
        TreeNode node = new TreeNode(nums[half]);
        node.left = generateBST(nums, start, half - 1);
        node.right = generateBST(nums, half + 1, end);
        return node;
    }

    public TreeNode sortedArrayToBST(int[] nums) {
        if(nums.length == 0) {
            return null;
        }
        return generateBST(nums, 0, nums.length - 1);
    }
}
