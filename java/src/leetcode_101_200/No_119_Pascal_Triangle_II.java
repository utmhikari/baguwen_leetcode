package leetcode_101_200;

import java.util.*;
/**
 * Created by 36191 on 2017/12/15
 * 119.读取杨辉三角序列特定行
 */

public class No_119_Pascal_Triangle_II {
    public List<Integer> getRow(int rowIndex) {
        ArrayList<Integer> nums = new ArrayList<>();
        if(rowIndex < 0) {
            return nums;
        }
        for(int i = 0; i <= rowIndex; i++) {
            nums.add(0);
        }
        nums.set(0, 1);
        for(int i = 1; i <= rowIndex; i++) {
            nums.set(i, 1);
            for(int j = i - 1; j >= 1; j--) {
                nums.set(j, nums.get(j) + nums.get(j - 1));
            }
        }
        return nums;
    }
}
