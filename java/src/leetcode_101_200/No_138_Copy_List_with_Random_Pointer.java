package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/18
 */

import java.util.*;
import utils.RandomListNode;

public class No_138_Copy_List_with_Random_Pointer {

    public RandomListNode copyRandomList(RandomListNode head) {
        if (head == null) { return null; }
        HashMap<String, Integer> addrNode = new HashMap<>(32);
        ArrayList<RandomListNode> nodes = new ArrayList<>();
        RandomListNode copy = new RandomListNode(head.label), p = head, c = copy;
        while (true) {
            addrNode.put(p.toString(), nodes.size());
            nodes.add(c);
            p = p.next;
            if (p == null) { break; }
            c.next = new RandomListNode(p.label);
            c = c.next;
        }
        p = head; c = copy;
        while (p != null) {
            if (p.random != null) {
                c.random = nodes.get(addrNode.get(p.random.toString()));
            }
            p = p.next;
            c = c.next;
        }
        return copy;
    }
}
