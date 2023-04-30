package leetcode_301_400;

import java.util.Stack;

/**
 * Created by 36191 on 2019/1/12
 */
public class No_331_Verify_Preorder_Serialization_of_a_Binary_Tree {
    public boolean isValidSerialization(String preorder) {
        if (preorder == null) {
            return false;
        }
        Stack<String> st = new Stack<>();
        String[] strs = preorder.split(",");
        for (int pos = 0; pos < strs.length; pos++) {
            String cur = strs[pos];
            while (cur.equals("#") && !st.isEmpty() && st.peek().equals(cur)) {
                st.pop();
                if (st.isEmpty()) { return false; }
                st.pop();
            }
            st.push(cur);
        }
        return st.size() == 1 && st.peek().equals("#");
    }
}
