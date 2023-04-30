package leetcode_101_200;

import utils.TreeLinkNode;
import java.util.*;
/**
 * Created by 36191 on 2017/12/13
 * 116.连树链表中右边节点
 */

public class No_116_Populating_Next_Right_Pointers_in_Each_Node {
    public void connect(LinkedList<TreeLinkNode> nodes) {
        int size = nodes.size();
        if(size != 0) {
            for(int i = 0; i < size - 1; i++) {
                nodes.get(i).next = nodes.get(i + 1);
            }
            for(int i = 0; i < size; i++) {
                if(nodes.get(i).left != null) {
                    nodes.add(nodes.get(i).left);
                }
                if(nodes.get(i).right != null) {
                    nodes.add(nodes.get(i).right);
                }
            }
            for(int i = 0; i < size; i++) {
                nodes.removeFirst();
            }
            connect(nodes);
        }
    }

    public void connect(TreeLinkNode root) {
        if(root != null) {
            LinkedList<TreeLinkNode> nodes = new LinkedList<>();
            nodes.add(root);
            connect(nodes);
        }
    }
}
