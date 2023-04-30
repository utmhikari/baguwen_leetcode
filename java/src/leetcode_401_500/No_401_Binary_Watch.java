package leetcode_401_500;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_401_Binary_Watch {
    private String getTime(List<Integer> indices) {
        int left = 0, right = 0;
        for (int i : indices) {
            if (i <= 3) { left += (1 << i); }
            else { right += (1 << (i - 4)); }
        }
        if (left > 11 || right > 59) { return ""; }
        StringBuilder sb = new StringBuilder();
        sb.append(left);
        sb.append(":");
        if (right == 0) { sb.append("00"); }
        else if (right < 10) { sb.append('0'); sb.append(right); }
        else { sb.append(right); }
        return sb.toString();
    }

    private void dfs(List<String> ans, List<Integer> indices, int start, int num, int max) {
        indices.add(start);
        int size = indices.size();
        if (size == num) {
            String s = getTime(indices);
            if (s.length() != 0) { ans.add(s); }
        } else {
            for (int i = start + 1; i <= max - num + size; i++) {
                dfs(ans, indices, i, num, max);
            }
        }
        indices.remove(indices.size() - 1);
    }

    public List<String> readBinaryWatch(int num) {
        List<String> ans = new ArrayList<String>();
        if (num == 0) { ans.add("0:00"); return ans; }
        List<Integer> indices = new ArrayList<>();
        for (int i = 0; i <= 10 - num; i++) { dfs(ans, indices, i, num, 10); }
        return ans;
    }
}
