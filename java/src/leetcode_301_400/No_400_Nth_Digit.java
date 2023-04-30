package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_400_Nth_Digit {
    public int findNthDigit(int n) {
        long sum = 0;
        int pow = 1;
        while (sum < n){
            long start = (long)Math.pow(10, pow - 1), end = (long)Math.pow(10, pow), acc =  (end - start) * pow;
            if (sum + acc >= n) {
                long diff = n - sum, num = diff / pow + start - 1;
                int rem = (int)(diff % pow);
                return rem > 0 ? String.valueOf(num + 1).charAt(rem - 1) - '0' : String.valueOf(num).charAt(pow - 1) - '0';
            }
            sum += acc;
            pow ++;
        }
        return 0;
    }
}
