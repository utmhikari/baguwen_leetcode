package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/20
 */

import utils.TreeNode;

import java.util.ArrayList;
import java.util.List;

public class No_236_Lowest_Common_Ancestor_of_a_Binary_Tree {

    private static boolean findP, findQ;

    private void traverse(List<ArrayList<TreeNode>> routes, ArrayList<TreeNode> route,
                             TreeNode root, TreeNode p, TreeNode q) {
        if (!(findP && findQ)) {
            route.add(root);
            if (!findP && root.val == p.val) { findP = true; routes.add(new ArrayList<>(route)); }
            else if (!findQ && root.val == q.val) { findQ = true; routes.add(new ArrayList<>(route)); }
            if (root.left != null) { traverse(routes, route, root.left, p, q); }
            if (root.right != null) { traverse(routes, route, root.right, p, q); }
            route.remove(route.size() - 1);
        }
    }

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        findP = false;
        findQ = false;
        List<ArrayList<TreeNode>> routes = new ArrayList<>();
        ArrayList<TreeNode> route = new ArrayList<>();
        traverse(routes, route, root, p, q);
        int l = Math.min(routes.get(0).size(), routes.get(1).size());
        for (int i = 1; i < l; i++) {
            if (routes.get(0).get(i).val != routes.get(1).get(i).val) {
                return routes.get(0).get(i - 1);
            }
        }
        return routes.get(0).get(l - 1);
    }
}
