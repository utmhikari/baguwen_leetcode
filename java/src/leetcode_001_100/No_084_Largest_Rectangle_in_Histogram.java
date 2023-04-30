package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/23
 * 084. 柱状图最大面积正方形
 */
import java.util.LinkedList;
import java.util.Stack;
public class No_084_Largest_Rectangle_in_Histogram {


    /* https://discuss.leetcode.com/topic/7599/o-n-stack-based-java-solution */
    public int largestRectangleArea(int[] height) {
        int len = height.length;
        Stack<Integer> s = new Stack<Integer>();
        int maxArea = 0;
        for(int i = 0; i <= len; i++){
            int h = (i == len ? 0 : height[i]);
            if(s.isEmpty() || h >= height[s.peek()]){
                s.push(i);
            }else{
                int tp = s.pop();
                maxArea = Math.max(maxArea, height[tp] * (s.isEmpty() ? i : i - 1 - s.peek())); // pop直到height更低为止
                i--;
            }
        }
        return maxArea;
    }

    /*
    public int largestRectangleArea(int[] heights, int startIndex, int endIndex) {
        if(startIndex > endIndex) {
            return 0;
        }
        if(startIndex == endIndex) {
            return heights[startIndex];
        }
        int maxArea = Integer.MIN_VALUE;
        int maxIndex = -1;
        for(int i = startIndex; i <= endIndex; i++) {
            if(heights[i] > maxArea) {
                maxArea = heights[i];
                maxIndex = i;
            }
        }
        int l = maxIndex;
        int r = maxIndex;
        int height = Math.min(heights[l], heights[r]);
        while(true) {
            int area = (r - l + 1) * height;
            if(area > maxArea) {
                maxArea = area;
            }
            if(l == startIndex && r == endIndex) {
                break;
            } else if(l == startIndex){
                r++;
                height = Math.min(height, heights[r]);
            } else if(r == endIndex) {
                l--;
                height = Math.min(height, heights[l]);
            } else {
                if(heights[r + 1] > heights[l - 1]) {
                    r++;
                    height = Math.min(height, heights[r]);
                } else {
                    l--;
                    height = Math.min(height, heights[l]);
                }
            }
        }
        return maxArea;
}

    public int largestRectangleArea(int[] heights) {
        int len = heights.length;
        if(len == 0) {
            return 0;
        } else if(len == 1){
            return heights[0];
        } else {
            int start = 0;
            int maxArea = 0;
            LinkedList<int[]> scales = new LinkedList<>();
            for(int i = 0; i < len; i++) {
                if(heights[i] == 0) {
                    scales.add(new int[] {start, i - 1});
                    start = i + 1;
                    if(start == len) {
                        break;
                    }
                }
            }
            if(start != len) {
                scales.add(new int[] {start, len - 1});
            }
            for(int[] scale : scales) {
                int area = largestRectangleArea(heights, scale[0], scale[1]);
                if(area > maxArea) {
                    maxArea = area;
                }
            }
            return maxArea;
        }
    }
    */

    public static void main(String[] args) {
        No_084_Largest_Rectangle_in_Histogram solution = new No_084_Largest_Rectangle_in_Histogram();
        System.out.println(solution.largestRectangleArea(new int[] {4,2,0,3,2,4,3,4}));
        System.out.println(solution.largestRectangleArea(new int[] {1, 1}));
        System.out.println(solution.largestRectangleArea(new int[] {0, 0, 0, 0}));
        System.out.println(solution.largestRectangleArea(new int[] {5, 4, 1, 2}));
    }
}
