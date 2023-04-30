package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/18
 * 067. 二进制加法
 */
import java.util.LinkedList;

public class No_067_Add_Binary {
    public String addLargeSmall(LinkedList<Integer> sum, String large, String small, int lenLarge, int lenSmall) {
        StringBuilder result = new StringBuilder();
        int countOne = 0;
        for(int i = 0; i < lenSmall; i++) {
            if(small.charAt(lenSmall - 1 - i) == '1') {
                countOne++;
            }
            if(large.charAt(lenLarge - 1 - i) == '1') {
                countOne++;
            }
            if(countOne < 2) {
                sum.add(countOne);
                countOne = 0;
            } else {
                sum.add(countOne - 2);
                countOne = 1;
            }
        }
        int remainLen = lenLarge - lenSmall;
        for(int i = 0; i < remainLen; i++) {
            if(large.charAt(remainLen - 1 - i) == '1') {
                countOne++;
            }
            if(countOne < 2) {
                sum.add(countOne);
                countOne = 0;
            } else {
                sum.add(countOne - 2);
                countOne = 1;
            }
        }
        if(countOne == 1) {
            sum.add(1);
        }
        while(sum.size() != 0) {
            result.append(sum.getLast());
            sum.removeLast();
        }
        return result.toString();
    }

    public String addBinary(String a, String b) {
        LinkedList<Integer> sum = new LinkedList<>();
        int lenA = a.length();
        int lenB = b.length();
        if(lenA > lenB) {
            return addLargeSmall(sum, a, b, lenA, lenB);
        } else {
            return addLargeSmall(sum, b, a, lenB, lenA);
        }
    }

    public static void main(String[] args) {
        No_067_Add_Binary solution = new No_067_Add_Binary();
        System.out.println(solution.addBinary("1111100001", "11111"));
    }
}
