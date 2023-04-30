package leetcode_301_400;

import java.util.HashMap;

/**
 * Created by 36191 on 2019/1/14
 */
public class No_357_Count_Numbers_with_Unique_Digits {
    
    public static HashMap<Integer, Integer> map = new HashMap<Integer, Integer>() {{
        put(0, 1); put(1, 10); put(2, 91);
    }};
    
    public int countNumbersWithUniqueDigits(int n) {
        if (n > 10) { return countNumbersWithUniqueDigits(10); }
        if (map.containsKey(n)) { return map.get(n); }
        return (12 - n) * countNumbersWithUniqueDigits(n - 1) - (11 - n) * countNumbersWithUniqueDigits(n - 2);
    }
    
    public static void main(String[] args) {
        No_357_Count_Numbers_with_Unique_Digits solution = new No_357_Count_Numbers_with_Unique_Digits();
        for (int i = 0; i < 10; i++) {
            System.out.println(solution.countNumbersWithUniqueDigits(i));
        }
    }
    
}
