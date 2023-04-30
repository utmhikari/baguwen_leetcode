package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */

public class No_165_Compare_Version_Numbers {

    public int compareVersion(String version1, String version2) {
        String[] vsp1 = version1.split("\\.");
        String[] vsp2 = version2.split("\\.");
        int i = 0;
        while (i < vsp1.length && i < vsp2.length) {
            int i1 = Integer.parseInt(vsp1[i]), i2 = Integer.parseInt(vsp2[i]);
            if (i1 < i2) { return -1; }
            else if (i1 > i2) { return 1; }
            else { ++i; }
        }
        if (i == vsp1.length && i == vsp2.length) { return 0; }
        else if (i == vsp1.length) {
            while (i < vsp2.length) {
                if (Integer.parseInt(vsp2[i]) > 0) { return -1; }
                else { ++i; }
            }
            return 0;
        } else {
            while (i < vsp1.length) {
                if (Integer.parseInt(vsp1[i]) > 0) { return 1; }
                else { ++i; }
            }
            return 0;
        }
    }

    public static void main(String[] args) {
        No_165_Compare_Version_Numbers solution = new No_165_Compare_Version_Numbers();
        String[][] inputs = new String[][]{
                new String[]{"0.1", "1.1"},
                new String[]{"1.0.1", "1"},
                new String[]{"7.5.2.4", "7.5.3"},
        };
        for (String[] input : inputs) {
            System.out.println(solution.compareVersion(input[0], input[1]));
        }
    }
}
