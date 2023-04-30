package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/1
 */
public class No_205_Isomorphic_Strings {
    public boolean isIsomorphic(String s, String t) {
        int[] kankeis = new int[128], kankeit = new int[128];
        for (int i = 0, l = s.length(); i < l; i++) {
            char cs = s.charAt(i), ct = t.charAt(i);
            if (kankeis[cs] == 0) { kankeis[cs] = ct; }
            else if (kankeis[cs] != ct) { return false; }
            if (kankeit[ct] == 0) { kankeit[ct] = cs; }
            else if (kankeit[ct] != cs) { return false; }
        }
        return true;
    }

    public static void main(String[] args) {
        No_205_Isomorphic_Strings solution = new No_205_Isomorphic_Strings();
    }
}
