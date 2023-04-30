package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_260_Single_Number_III {
    public int[] singleNumber(int[] nums) {
        int xor = 0;
        for (int i : nums) {
            xor ^= i;
        }
        int bitMask = ((xor - 1) & xor) ^ xor;
        int a = 0, b = 0;
        for (int i : nums) {
            if ((i & bitMask) == 0) {
                a ^= i;
            } else {
                b ^= i;
            }
        }
        return new int[] {a, b};
    }
}
