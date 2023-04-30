package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/15
 * 050. x的n次方
 */

public class No_050_Pow_x_n {

    public double myPow(double x, int n) {
        double val = 1;
        if(n == 0) {
            return val;
        }
        double half = myPow(x, n / 2);
        if(n % 2 == 0) {
            return half * half;
        } else {
            if(n > 0) {
                return x * half * half;
            } else {
                return (1 / x) * half * half;
            }
        }
    }

    public static void main(String[] args) {
        No_050_Pow_x_n solution = new No_050_Pow_x_n();
        System.out.println(solution.myPow(3.14, 2));
        System.out.println(solution.myPow(3.14, 5));
        System.out.println(solution.myPow(3.14, -3));
    }
}
