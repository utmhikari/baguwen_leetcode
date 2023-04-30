package leetcode_401_500;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_414_Third_Maximum_Number {
    public int thirdMax(int[] nums) {
        Integer max1 = null, max2 = null, max3 = null;
        for (Integer n : nums) {
            if (!n.equals(max1) && !n.equals(max2) && !n.equals(max3)) {
                if (max1 == null || n > max1) { max3 = max2; max2 = max1; max1 = n; }
                else if (max2 == null || n > max2) { max3 = max2; max2 = n; }
                else if (max3 == null || n > max3) { max3 = n; }
            }
        }
        return max3 == null ? max1 : max3;
    }
}
