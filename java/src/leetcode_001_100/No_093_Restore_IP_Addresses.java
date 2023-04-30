package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/6
 * 093. 识别IP地址
 */

import java.util.List;
import java.util.ArrayList;
public class No_093_Restore_IP_Addresses {

    public List<String> restoreIpAddresses(String s) {
        List<String> res = new ArrayList<>();
        int len = s.length();
        for(int i = 1; i < 4 && i < len - 2; i++){
            for(int j = i + 1; j < i + 4 && j < len - 1; j++){
                for(int k = j + 1; k < j + 4 && k < len; k++){
                    String s1 = s.substring(0, i), s2 = s.substring(i, j), s3 = s.substring(j, k), s4 = s.substring(k, len);
                    if(isValid(s1) && isValid(s2) && isValid(s3) && isValid(s4)){
                        res.add(s1 + "." + s2 + "." + s3 + "." + s4);
                    }
                }
            }
        }
        return res;
    }
    public boolean isValid(String s){
        return s.length() <= 3 && s.length() != 0 && (s.charAt(0) != '0' || s.length() <= 1) && Integer.parseInt(s) <= 255;
    }

    public void test(String s) {
        List<String> answers = restoreIpAddresses(s);
        for(String ans : answers) {
            System.out.println(ans);
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_093_Restore_IP_Addresses solution = new No_093_Restore_IP_Addresses();
        solution.test("25525511135");
    }
}
