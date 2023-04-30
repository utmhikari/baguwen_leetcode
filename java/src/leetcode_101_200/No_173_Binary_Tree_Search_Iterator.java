package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/25
 */

import utils.*;
import java.util.*;

public class No_173_Binary_Tree_Search_Iterator {

    class BSTIterator {

        private ArrayList<Integer> v;
        private int idx;

        private void pre(TreeNode node) {
            if (node != null) {
                pre(node.left);
                v.add(node.val);
                pre(node.right);
            }
        }

        public BSTIterator(TreeNode root) {
            v = new ArrayList<>();
            idx = -1;
            pre(root);
        }

        /** @return whether we have a next smallest number */
        public boolean hasNext() {
            return idx + 1 < v.size();
        }

        /** @return the next smallest number */
        public int next() {
            idx += 1;
            return v.get(idx);
        }
    }

    public static void main(String[] args) {
        No_173_Binary_Tree_Search_Iterator solution = new No_173_Binary_Tree_Search_Iterator();
    }
}
