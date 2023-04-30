package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/4.
 * 回文数字（不得开其它空间）
 */
public class No_009_Palindrome_Number {

    public boolean isPalindrome(int x) {
        if(x < 0) {
            return false;
        }
        int _x = x;
        int value = 0;
        while(_x != 0) {
            int last = _x % 10;
            int temp = 10 * value + last;
            if((temp - last) / 10 != value) {
                return false;
            }
            value = temp;
            _x /= 10;
        }
        return value == x;
    }

    public static void main(String[] args) {
        No_009_Palindrome_Number solution = new No_009_Palindrome_Number();
        System.out.println(solution.isPalindrome(123));
        System.out.println(solution.isPalindrome(121));
        System.out.println(solution.isPalindrome(-123));
        System.out.println(solution.isPalindrome(-121));
    }
}
