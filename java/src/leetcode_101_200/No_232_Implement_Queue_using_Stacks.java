package leetcode_101_200;

import java.util.Stack;

/**
 * Created by 36191 on 2018/10/6
 */
public class No_232_Implement_Queue_using_Stacks {
    private class MyQueue {

        private Stack<Integer> q;

        /** Initialize your data structure here. */
        public MyQueue() {
            q = new Stack<>();
        }

        /** Push element x to the back of queue. */
        public void push(int x) {
            q.push(x);
        }

        /** Removes the element from in front of queue and returns that element. */
        public int pop() {
            if (q.empty()) { return Integer.MIN_VALUE; }
            int val = q.get(0);
            q.remove(0);
            return val;
        }

        /** Get the front element. */
        public int peek() {
            return q.empty() ? Integer.MIN_VALUE : q.get(0);
        }

        /** Returns whether the queue is empty. */
        public boolean empty() {
            return q.empty();
        }
    }

    public MyQueue newQueue() {
        return new MyQueue();
    }

    public static void main(String[] args) {
        No_232_Implement_Queue_using_Stacks solution = new No_232_Implement_Queue_using_Stacks();
        MyQueue queue = solution.newQueue();
        queue.push(1);
        queue.push(2);
        System.out.println(queue.peek());
        System.out.println(queue.pop());
        System.out.println(queue.empty());
    }
}
