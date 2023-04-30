package utils;

import java.util.*;

/**
 * Created by 36191 on 2018/8/26
 */


public class UndirectedGraphNode {
    public  int label;
    public List<UndirectedGraphNode> neighbors;
    public UndirectedGraphNode(int x) { label = x; neighbors = new ArrayList<UndirectedGraphNode>(); }
}
