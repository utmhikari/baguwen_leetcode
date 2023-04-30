package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/22
 */
public class No_278_First_Bad_Version {

    private boolean isBadVersion(int m) {
        return false;
    }

    private int find(int l, int r) {
        int m = l + (r - l) / 2;
        if (m == l) { return isBadVersion(m) ? m : m + 1; }
        return isBadVersion(m) ? find(l, m) : find(m, r);
    }

    public int firstBadVersion(int n) {
        return find(1, n);
    }
}
