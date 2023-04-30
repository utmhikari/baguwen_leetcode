package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/9
 */
public class No_393_Valid_UTF8_Validation {

    private static int cur;

    public boolean validUtf8(int[] data) {
        int len = data.length;
        cur = 0;
        while (true) {
            int head = data[cur];
            if ((head & 128) != 0) {
                // 1 byte, must be true
                if ((head & 224) == 192) {
                    // 2 bytes
                    if (cur + 2 > len) { return false; }
                    if ((data[++cur] & 192) != 128) { break; }
                } else if ((head & 240) == 224) {
                    // 3 bytes
                    if (cur + 3 > len) { return false; }
                    if ((data[++cur] & 192) != 128 || (data[++cur] & 192) != 128) { break; }
                } else if ((head & 248) == 240) {
                    // 4 bytes
                    if (cur + 4 > len) { return false; }
                    if ((data[++cur] & 192) != 128 || (data[++cur] & 192) != 128 || (data[++cur] & 192) != 128) { break; }
                } else {
                    break;
                }
            }
            if (++cur == len) { return true; }
        }
        return false;
    }
}
