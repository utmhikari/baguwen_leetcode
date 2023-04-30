package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/3.
 * 007.整数位数反置
 */
public class No_007_Reverse_Integer {

    public int reverse(int x) {
        int value = 0;
        while(x != 0) {
            int last = x % 10;
            int temp = 10 * value + last;
            // 溢出
            if((temp - last) / 10 != value) {
                return 0;
            }
            value = temp;
            x /= 10;
        }
        return value;
    }

    public static void main(String[] args) {
        No_007_Reverse_Integer solution = new No_007_Reverse_Integer();
        System.out.println(solution.reverse(123));
        System.out.println(solution.reverse(-123));
    }
}
