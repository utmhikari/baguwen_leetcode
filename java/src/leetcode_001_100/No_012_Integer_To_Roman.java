package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/5.
 * 012.整数到罗马数字转换
 */
public class No_012_Integer_To_Roman {

    public String intToRoman(int num) {
        StringBuilder roman = new StringBuilder();
        String[] ones = new String[]{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
        String[] tens = new String[]{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
        String[] hundreds = new String[]{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
        String[] thousands = new String[]{"", "M", "MM", "MMM"};
        roman.append(thousands[(num / 1000) % 10]);
        roman.append(hundreds[(num / 100) % 10]);
        roman.append(tens[(num / 10) % 10]);
        roman.append(ones[num % 10]);
        return roman.toString();
    }

    public void output(int num) {
        System.out.println(intToRoman(num));
    }

    public static void main(String[] args) {
        No_012_Integer_To_Roman solution = new No_012_Integer_To_Roman();
        solution.output(1);
        solution.output(4);
        solution.output(10);
        solution.output(19);
        solution.output(20);
        solution.output(45);
        solution.output(100);
        solution.output(499);
        solution.output(1000);
        solution.output(1980);
        solution.output(2999);
    }
}
