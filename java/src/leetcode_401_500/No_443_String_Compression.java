package leetcode_401_500;

/**
 * Created by 36191 on 2019/2/24
 */
public class No_443_String_Compression {
    public int compress(char[] chars) {
        if (chars == null) { return 0; }
        int cur = 0;
        for (int i = 0; i < chars.length; i++) {
            char c = chars[i];
            int cnt = 1;
            while ( (i + 1) < chars.length && chars[i] == chars[i + 1]) { ++i; ++cnt; }
            chars[cur++] = c;
            if (cnt != 1) {
                for (char d : String.valueOf(cnt).toCharArray()) {
                    chars[cur++] = d;
                }
            }
        }
        return cur;
    }
}
