package utils;

import java.util.List;

/**
 * Created by 36191 on 2019/2/12
 */
public class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int val,List<Node> children) {
        this.val = val;
        this.children = children;
    }
}
