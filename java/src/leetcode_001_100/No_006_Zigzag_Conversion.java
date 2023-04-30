package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/4.
 * 006.Zigzag输入转换
 * 例子：PAYPALISHIRING, 3
 * 转换成：PAHNAPLSIIGYIR
 * 因为输入如下
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 */
import java.util.ArrayList;
import java.util.HashMap;

public class No_006_Zigzag_Conversion {

    public String convert(String s, int numRows) {
        if(numRows <= 1) {
            return s;
        }
        HashMap<Integer, ArrayList<Character>> charMap = new HashMap<>();
        for(int i = 0; i < numRows; i++) {
            charMap.put(i, new ArrayList<Character>());
        }
        int len = s.length();
        for(int i = 0; i < len; i++) {
            char c = s.charAt(i);
            int modDouble = i % (2 * (numRows - 1));
            int modSingle = i % (numRows - 1);
            if(modDouble == 0) {
                charMap.get(0).add(c);
            } else if(modSingle == 0) {
                charMap.get(numRows - 1).add(c);
            } else if(modDouble == modSingle){
                charMap.get(modSingle).add(c);
            } else {
                charMap.get(numRows - 1 - modSingle).add(c);
            }
        }
        StringBuilder ans = new StringBuilder();
        for(int i = 0; i < numRows; i++) {
            for(char j : charMap.get(i)) {
                ans.append(j);
            }
        }
        return ans.toString();
    }

    public static void main(String[] args) {
        No_006_Zigzag_Conversion solution = new No_006_Zigzag_Conversion();
        System.out.println(solution.convert("PAYPALISHIRING", 3));
        System.out.println(solution.convert("PAY", 1));
    }
}
