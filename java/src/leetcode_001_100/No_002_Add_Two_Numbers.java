package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/2.
 * 002.两个链表相加
 */
public class No_002_Add_Two_Numbers {
    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    /**
     * 相加两个链表
     * @param l1
     * @param l2
     * @return
     */
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode first = new ListNode(-1);
        ListNode sum = first;
        int add = 0;
        while(l1 != null || l2 != null) {
            int value = add;
            if(l1 != null) {
                value += l1.val;
                l1 = l1.next;
            }
            if(l2 != null) {
                value += l2.val;
                l2 = l2.next;
            }
            if(value > 9) {
                value -= 10;
                add = 1;
            } else {
                add = 0;
            }
            sum.next = new ListNode(value);
            sum = sum.next;
        }
        if(add == 1) {
            sum.next = new ListNode(1);
        }
        return first.next;
    }

    public void test() {
        ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);
        ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);
        ListNode sum = addTwoNumbers(l1, l2);
        while(sum != null) {
            System.out.print(String.valueOf(sum.val) + ", ");
            sum = sum.next;
        }
    }

    public static void main(String[] args) {
        No_002_Add_Two_Numbers solution = new No_002_Add_Two_Numbers();
        solution.test();
    }
}
