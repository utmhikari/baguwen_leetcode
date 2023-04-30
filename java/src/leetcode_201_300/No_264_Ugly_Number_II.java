package leetcode_201_300;

import java.util.ArrayList;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_264_Ugly_Number_II {

    public int nthUglyNumber(int n) {
        if (n == 1) { return 1; }
        int two = 0, three = 0, five = 0;
        ArrayList<Integer> ugly = new ArrayList<Integer>() {{ add(1); }};
        while (ugly.size() < n) {
            int next = Math.min(ugly.get(two) * 2, Math.min(ugly.get(three) * 3, ugly.get(five) * 5));
            ugly.add(next);
            if (ugly.get(two) * 2 == next) { ++two; }
            if (ugly.get(three) * 3 == next) { ++three; }
            if (ugly.get(five) * 5 == next) { ++five; }
        }
        return ugly.get(ugly.size() - 1);
    }

    public static void main(String[] args) {

    }
}
