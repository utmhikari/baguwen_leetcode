package leetcode_101_200;

import java.util.*;
/**
 * Created by 36191 on 2017/12/15
 * 120.寻找三角序列中最低权重和
 */

public class No_120_Triangle {

    private int minValue = Integer.MAX_VALUE;

    private void traverse(List<List<Integer>> triangle, int[] sums, int rowIndex) {
        for(int i = triangle.get(rowIndex).size() - 1; i >= 0; i--) {
            int tempSum = triangle.get(rowIndex).get(i);
            if(i == triangle.get(rowIndex).size() - 1 && i != 0) {
                tempSum += sums[i - 1];
            } else if(i > 0) {
                tempSum += Math.min(sums[i - 1], sums[i]);
            } else {
                tempSum += sums[i];
            }
            sums[i] = tempSum;
            if(rowIndex == triangle.size() - 1) {
                if(sums[i] < minValue) {
                    minValue = sums[i];
                }
            }
        }
    }

    public int minimumTotal(List<List<Integer>> triangle) {
        if(triangle == null) {
            return 0;
        } else if(triangle.size() == 1) {
            return triangle.get(0).get(0);
        }
        int[] sums = new int[triangle.get(triangle.size() - 1).size()];
        minValue = Integer.MAX_VALUE;
        for(int i = 0; i < triangle.size(); i++) {
            traverse(triangle, sums, i);
        }
        return minValue;
    }
}
