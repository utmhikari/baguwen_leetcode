package leetcode_201_300;

import java.util.Stack;

/**
 * Created by 36191 on 2018/10/20
 */
public class No_232_Implement_Queue_Using_Stacks {
    class MyQueue {

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
}
