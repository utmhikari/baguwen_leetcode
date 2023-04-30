package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/9.
 * 029.整数相除
 */

public class No_029_Divide_Two_Integers {

    public int divide(int dividend, int divisor) {
        //Reduce the problem to positive long integer to make it easier.
        //Use long to avoid integer overflow cases.
        int sign = 1;
        if ((dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0))
            sign = -1;
        long ldividend = Math.abs((long) dividend);
        long ldivisor = Math.abs((long) divisor);

        //Take care the edge cases.
        if (ldivisor == 0) return Integer.MAX_VALUE;
        if ((ldividend == 0) || (ldividend < ldivisor))	return 0;

        long lans = ldivide(ldividend, ldivisor);

        int ans;
        if (lans > Integer.MAX_VALUE){ //Handle overflow.
            ans = (sign == 1)? Integer.MAX_VALUE : Integer.MIN_VALUE;
        } else {
            ans = (int) (sign * lans);
        }
        return ans;
    }

    private long ldivide(long ldividend, long ldivisor) {
        // Recursion exit condition
        if (ldividend < ldivisor) return 0;

        //  Find the largest multiple so that (divisor * multiple <= dividend),
        //  whereas we are moving with stride 1, 2, 4, 8, 16...2^n for performance reason.
        //  Think this as a binary search.
        long sum = ldivisor;
        long multiple = 1;
        while ((sum+sum) <= ldividend) {
            sum += sum;
            multiple += multiple;
        }
        //Look for additional value for the multiple from the reminder (dividend - sum) recursively.
        return multiple + ldivide(ldividend - sum, ldivisor);
    }

    /*
    public int divide(int dividend, int divisor) {
        if(divisor == 0 || (dividend == Integer.MIN_VALUE && divisor == -1)) {
            return Integer.MAX_VALUE;
        }
        if(dividend == 0 || divisor == Integer.MIN_VALUE) {
            return 0;
        }
        if(divisor == 1) {
            return dividend;
        }
        if(divisor == -1) {
            return 0 - dividend;
        }
        int val = 0;
        int temp;
        int add;
        if(dividend == Integer.MIN_VALUE) {
            temp = Integer.MAX_VALUE;
            add = 1;
        } else {
            temp = Math.abs(dividend);
            add = 0;
        }
        int _temp = Math.abs(divisor);
        boolean pos;
        if(dividend < 0 && divisor < 0 || dividend > 0 && divisor > 0) {
            pos = true;
        } else {
            pos = false;
        }
        int multiple = 1; // 步数按2的n次幂
        while(temp >= _temp) {
            if(pos) {
                val++;
            } else {
                val--;
            }
            _temp += _temp;
            multiple += multiple;
            if(add == 1) {
                _temp += 1;
                add = 0;
            }
        }
        return val;
    }
    */

    public static void main(String[] args) {
        No_029_Divide_Two_Integers solution = new No_029_Divide_Two_Integers();
        System.out.println(solution.divide(Integer.MIN_VALUE, -2));
        System.out.println(solution.divide(Integer.MIN_VALUE, 2));
        System.out.println(solution.divide(Integer.MAX_VALUE, -2));
        System.out.println(solution.divide(Integer.MAX_VALUE, Integer.MIN_VALUE));
        System.out.println(solution.divide(Integer.MIN_VALUE, Integer.MAX_VALUE));
        System.out.println(solution.divide(6, -2));
    }
}
