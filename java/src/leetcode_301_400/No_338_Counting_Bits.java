package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_338_Counting_Bits {
    public int[] countBits(int num) {
        if (num == 0) { return new int[]{0}; }
        int[] nums = new int[num + 1];
        nums[0] = 0;
        int idx = 0;
        while (true) {
            int nextIdx = ((idx + 1) << 1) - 1;
            if (nextIdx < num) {
                for (int i = 0; i <= idx; i++) { nums[i + idx + 1] = 1 + nums[i]; }
                idx = nextIdx;
            } else {
                for (int i = idx + 1; i <= num; i++) { nums[i] = 1 + nums[i - 1 - idx]; }
                break;
            }
        }
        return nums;
    }
    
    public static void main(String[] args) {
        No_338_Counting_Bits solution = new No_338_Counting_Bits();
        for (int i = 0; i <= 16; i++) {
            int[] nums = solution.countBits(i);
            System.out.println("Number " + i + ": ");
            for (int n : nums) { System.out.print(n + ", "); }
            System.out.println();
        }
    }
    
}
