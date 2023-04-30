package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/3
 */

import java.util.Iterator;
import java.util.LinkedList;
import java.util.Queue;

public class No_225_Implement_Stack_Using_Queues {
    private class MyStack {
        private Queue<Integer> s;

        /** Initialize your data structure here. */
        public MyStack() {
            s = new LinkedList<>();
        }

        /** Push element x onto stack. */
        public void push(int x) {
            s.offer(x);
        }

        /** Removes the element on top of the stack and returns that element. */
        public int pop() {
            int v = Integer.MIN_VALUE;
            Iterator<Integer> iter = s.iterator();
            while (iter.hasNext()) { v = iter.next(); }
            iter.remove();
            return v;
        }

        /** Get the top element. */
        public int top() {
            int v = Integer.MIN_VALUE;
            Iterator<Integer> iter = s.iterator();
            while (iter.hasNext()) { v = iter.next(); }
            return v;
        }

        /** Returns whether the stack is empty. */
        public boolean empty() {
            return s.isEmpty();
        }
    }

    public MyStack stack() {
        return new MyStack();
    }

    public static void main(String[] args) {
        No_225_Implement_Stack_Using_Queues solution = new No_225_Implement_Stack_Using_Queues();
        MyStack stack = solution.stack();
        stack.push(1);
        stack.push(2);
        System.out.println(stack.top());
        System.out.println(stack.pop());
        System.out.println(stack.empty());
    }
}
