package leetcode_301_400;

import java.util.Iterator;
import java.util.List;
import java.util.Stack;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_341_Flatten_Nested_List_Iterator {

    public interface NestedInteger {
        /**
         * @return true if this NestedInteger holds a single integer, rather than a nested list.
         */
        public boolean isInteger();

        /**
         * @return the single integer that this NestedInteger holds, if it holds a single integer
         *           or null if this NestedInteger holds a nested list
         */
        public Integer getInteger();

        /**
         * 
         * @return the nested list that this NestedInteger holds, if it holds a nested list
         *          or null if this NestedInteger holds a single integer
         */
        public List<NestedInteger> getList();
    }
    
    public class NestedIterator implements Iterator<Integer> {
        Stack<NestedInteger> stack = new Stack<>();
        
        public NestedIterator(List<NestedInteger> nestedList) {
            for(int i = nestedList.size() - 1; i >= 0; i--) {
                stack.push(nestedList.get(i));
            }
        }

        @Override
        public Integer next() {
            return stack.pop().getInteger();
        }

        @Override
        public boolean hasNext() {
            while(!stack.isEmpty()) {
                NestedInteger next = stack.peek();
                if (next.isInteger()) { return true; }
                stack.pop();
                for(int i = next.getList().size() - 1; i >= 0; i--) {
                    stack.push(next.getList().get(i));
                }
            }
            return false;
        }
    }
}
