package leetcode_101_200;

import java.util.*;
/**
 * Created by 36191 on 2017/12/15
 * 118.生成杨辉三角序列
 */

public class No_118_Pascal_Triangle {
    public List<List<Integer>> generate(int numRows) {
        LinkedList<List<Integer>> ans = new LinkedList<>();
        if(numRows >= 1) {
            ArrayList<Integer> first = new ArrayList<Integer>(){{ add(1); }};
            ans.add(first);
            for(int i = 1; i < numRows; i++) {
                ArrayList<Integer> row = new ArrayList<>();
                row.add(1);
                for(int j = 0; j < ans.get(i - 1).size() - 1; j++) {
                    row.add(ans.get(i - 1).get(j) + ans.get(i - 1).get(j + 1));
                }
                row.add(1);
                ans.add(row);
            }
        }
        return ans;
    }
}
