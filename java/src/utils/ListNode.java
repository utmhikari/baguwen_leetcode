package utils;

/**
 * Created by 36191 on 2017/12/5
 * 链表
 */

public class ListNode {
    public int val;
    public ListNode next;
    public ListNode(int x) { val = x; }

    private void addNodes(ListNode node, int[] vals, int idx) {
        if (idx >= vals.length) { return; }
        node.next = new ListNode(vals[idx]);
        addNodes(node.next, vals, idx + 1);
    }

    public ListNode(int[] vals) {
        if (vals == null) { this.val = Integer.MIN_VALUE; }
        else {
            this.val = vals[0];
            addNodes(this, vals, 1);
        }
    }

    public void output() {
        StringBuilder str = new StringBuilder();
        ListNode cur = this;
        while (true) {
            str.append(cur.val);
            if (cur.next == null) { break; }
            str.append("->");
            cur = cur.next;
        }
        System.out.println(str.toString());
    }
}
