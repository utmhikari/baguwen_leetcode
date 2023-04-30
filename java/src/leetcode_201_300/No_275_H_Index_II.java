package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/22
 */
public class No_275_H_Index_II {
    public int hIndex(int[] citations) {
        int l = citations.length, i = 0, h = l;
        while (i < h){
            int m = i + h >> 1;
            if (citations[m] == l - m) { return l - m; }
            else if (citations[m] < l - m) { i = m + 1; }
            else { h = m; }
        }
        return l - i;
    }
}
