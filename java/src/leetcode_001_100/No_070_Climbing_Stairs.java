package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 070. 爬楼梯
 */

public class No_070_Climbing_Stairs {

    public int climbStairs(int n) {
        if(n <= 1) {
            return 1;
        } else {
            int f1 = 1;
            int f2 = 1;
            int f3 = 1;
            int i = 2;
            while(i <= n) {
                f3 = f1 + f2;
                f1 = f2;
                f2 = f3;
                i++;
            }
            return f3;
        }
    }

    public static void main(String[] args) {
        No_070_Climbing_Stairs solution = new No_070_Climbing_Stairs();
        System.out.println(solution.climbStairs(44));
    }
}
