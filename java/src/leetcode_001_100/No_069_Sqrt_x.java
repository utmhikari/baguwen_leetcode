package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 069. 开方
 */

public class No_069_Sqrt_x {
    public int mySqrt(int x) {
        if(x < 0) {
            return -1;
        } else if(x == 0 || x == 1) {
            return x;
        } else {
            double y = (double)(x / 2);
            for(int i = 0; i < 20; i++) {
                y = (y + x / y) / 2;
            }
            return (int)y;
        }
    }

    public static void main(String[] args) {
        No_069_Sqrt_x solution = new No_069_Sqrt_x();
        System.out.println(solution.mySqrt(0));
        System.out.println(solution.mySqrt(1));
        System.out.println(solution.mySqrt(2));
        System.out.println(solution.mySqrt(3));
        System.out.println(solution.mySqrt(31));
        System.out.println(solution.mySqrt(49));
        System.out.println(solution.mySqrt(Integer.MAX_VALUE));
    }
}
