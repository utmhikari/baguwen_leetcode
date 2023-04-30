package leetcode_401_500;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_434_Number_of_Segments_in_a_String {
    public int countSegments(String s) {
        char[] chars = s.toCharArray();
        int cnt = 0, left = 0, len = chars.length, right = len - 1;
        while (left < len && chars[left] == ' ') { ++left; }
        while (right >= 0 && chars[right] == ' ') { --right; }
        if (left > right) { return cnt; }
        while (true) {
            if (left > right) { return cnt + 1; }
            if (chars[left] != ' ') { ++left; }
            else {
                ++cnt;
                while (left <= right && chars[left] == ' ') { ++left; }
                if (left > right) { return cnt; }
            }
        }
    }
}
