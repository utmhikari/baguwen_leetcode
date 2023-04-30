package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/26
 */

import utils.*;

import java.util.*;

public class No_133_Clone_Graph {

    private UndirectedGraphNode clone(UndirectedGraphNode node, HashMap<Integer, UndirectedGraphNode> map) {
        if (node == null) {
            return null;
        }
        UndirectedGraphNode cNode;
        if (!map.containsKey(node.label)) {
            cNode = new UndirectedGraphNode(node.label);
            map.put(cNode.label, cNode);
            for (UndirectedGraphNode n : node.neighbors) {
                cNode.neighbors.add(clone(n, map));
            }
            return cNode;
        } else {
            return map.get(node.label);
        }
    }

    public UndirectedGraphNode cloneGraph(UndirectedGraphNode node) {
        if (node == null) {
            return null;
        }
        HashMap<Integer, UndirectedGraphNode> map = new HashMap<>();
        return clone(node, map);
    }

    public static void main(String[] args) {
        UndirectedGraphNode zero = new UndirectedGraphNode(0);
        UndirectedGraphNode one = new UndirectedGraphNode(1);
        UndirectedGraphNode two = new UndirectedGraphNode(2);
        zero.neighbors.add(one);
        zero.neighbors.add(two);
        one.neighbors.add(two);
        two.neighbors.add(two);
        UndirectedGraphNode clone = new No_133_Clone_Graph().cloneGraph(zero);
    }
}
