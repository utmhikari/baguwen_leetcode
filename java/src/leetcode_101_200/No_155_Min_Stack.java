package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 *
 */

public class No_155_Min_Stack {

    /**
     * Your MinStack object will be instantiated and called as such:
     * MinStack obj = new MinStack();
     * obj.push(x);
     * obj.pop();
     * int param_3 = obj.top();
     * int param_4 = obj.getMin();
     */

    static class MinStack {

        private class Node {
            int val;
            Node next;
            Node(int v) { val = v; }
        }

        private Node top;

        /** initialize your data structure here. */
        public MinStack() {
            top = null;
        }

        public void push(int x) {
            Node newTop = new Node(x);
            newTop.next = top;
            top = newTop;
        }

        public void pop() {
            if (top != null) {
                Node temp = top.next;
                top.next = null;
                top = temp;
            }
        }

        public int top() {
            if (top == null) { return Integer.MAX_VALUE; }
            return top.val;
        }

        public int getMin() {
            int min = Integer.MAX_VALUE;
            Node nd = top;
            while (nd != null) {
                min = Math.min(nd.val, min);
                nd = nd.next;
            }
            return min;
        }
    }

    public static void main(String[] args) {
        MinStack minStack = new MinStack();
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);
        minStack.getMin();
        minStack.pop();
        minStack.top();
        minStack.getMin();
    }

}

