package leetcode_401_500;

import java.util.ArrayList;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_405_Convert_a_Number_to_Hexadecimal {
    public String toHex(int num) {
        if (num == 0) { return "0"; }
        String hex = "0123456789abcdef";
        ArrayList<Character> chars = new ArrayList<>();
        for (int i = 0; i < 8; i++) {
            int idx = 0;
            for (int j = 0; j < 4; j++) {
                if ((num & num - 1) == num - 1) { idx += (1 << j); }
                num >>= 1;
                if (num == 0) { break; }
            }
            chars.add(hex.charAt(idx));
            if (num == 0) { break; }
        }
        StringBuilder sb = new StringBuilder();
        for (int i = chars.size() - 1; i >= 0; i--) { sb.append(chars.get(i)); }
        return sb.toString();
    }
}
