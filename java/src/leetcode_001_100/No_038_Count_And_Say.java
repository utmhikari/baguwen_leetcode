package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/12
 * 038. 数字符串的数字
 */

public class No_038_Count_And_Say {

    public String transform(String lastStr) {
        int len = lastStr.length();
        char temp = lastStr.charAt(0);
        int count = 1;
        StringBuilder sb = new StringBuilder();
        for(int i = 1; i < len; i++) {
            char c = lastStr.charAt(i);
            if(c != temp) {
                sb.append(count);
                sb.append((int)temp - 48);
                temp = c;
                count = 1;
            } else {
                count++;
            }
        }
        sb.append(count);
        sb.append((int)temp - 48);
        return sb.toString();
    }

    public String countAndSay(int n) {
        return n == 1 ? "1" : transform(countAndSay(n - 1));
    }

    public static void main(String[] args) {

    }
}
