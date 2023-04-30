package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/5.
 * 011.装最多水的容器
 */
public class No_011_Container_With_Most_Water {

    public int maxArea(int[] height) {
        int maxValue = 0;
        int len = height.length;
        int i = 0;
        int j = len - 1;
        while(j > i) {
            int a = j - i, b, area;
            if (height[i] < height[j]) {
                b = height[i];
                i++;
            } else {
                b = height[j];
                j--;
            }
            area = a * b;
            if (area >= maxValue) {
                maxValue = area;
            }
        }
        return maxValue;
    }

    public static void main(String[] args) {
        No_011_Container_With_Most_Water solution = new No_011_Container_With_Most_Water();
        System.out.println(solution.maxArea(new int[]{2, 1}));
        System.out.println(solution.maxArea(new int[]{2, 3, 4, 5, 18, 17, 6}));
    }
}
