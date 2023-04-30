package leetcode_301_400;

/**
 * Created by 36191 on 2018/12/1
 */
public class No_319_Bulb_Switcher {

    public int bulbSwitch(int n) {
        int ons = 0;
        long x = 1;
        while (x * x <= (long)n) { x++; ons++; }
        return ons;
    }

    public static void main(String[] args) {

    }
}
