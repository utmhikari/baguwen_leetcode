package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/4
 * 091. 解码方法总数
 */

public class No_091_Decode_Ways {

    public int numDecodings(char[] chars, int endIndex, int[] answers) {
        if(endIndex <= 0) {
            return 1;
        } else {
            if(answers[endIndex] != 0) {
                return answers[endIndex];
            }
            int sum = 0;
            if(chars[endIndex] == '0') {
                if(!(chars[endIndex - 1] != '1' && chars[endIndex - 1] != '2')) {
                    sum += numDecodings(chars, endIndex - 2, answers);
                }
            } else {
                sum += numDecodings(chars, endIndex - 1, answers);
                if(chars[endIndex - 1] == '1' || (chars[endIndex] <= '6' && chars[endIndex - 1] == '2')) {
                    sum += numDecodings(chars, endIndex - 2, answers);
                }
            }
            answers[endIndex] = sum;
            return sum;
        }
    }

    public int numDecodings(String s) {
        if(s.length() == 0) {
            return 0;
        } else if(s.length() == 1) {
            if(s.equals("0")) {
                return 0;
            }
            return 1;
        }
        char[] chars = s.toCharArray();
        int[] answers = new int[chars.length];
        if(chars[0] == '0') {
            return 0;
        }
        answers[0] = 1;
        return numDecodings(chars, chars.length - 1, answers);
    }

    public static void main(String[] args) {
        No_091_Decode_Ways solution = new No_091_Decode_Ways();
        System.out.println(solution.numDecodings("12"));
        System.out.println(solution.numDecodings("26"));
    }
}
