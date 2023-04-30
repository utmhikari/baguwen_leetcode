package leetcode_401_500;

import utils.Node;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_429_Nary_Tree_Level_Order_Traversal {

    private void traverse(List<List<Integer>> lvls, Node root, int depth) {
        if (root != null) {
            if (lvls.size() < depth + 1) {
                while (lvls.size() < depth + 1) {
                    lvls.add(new ArrayList<>());
                }
            }
            lvls.get(depth).add(root.val);
            if (root.children != null) {
                for (Node n : root.children) {
                    traverse(lvls, n, depth + 1);
                }
            }
        }
    }

    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> ans = new ArrayList<>();
        traverse(ans, root, 0);
        return ans;
    }
}
