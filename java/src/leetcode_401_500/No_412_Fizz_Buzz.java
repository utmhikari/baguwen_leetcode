package leetcode_401_500;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_412_Fizz_Buzz {
    public List<String> fizzBuzz(int n) {
        List<String> ans = new ArrayList<>();
        String[] fb = new String[] {"Fizz", "Buzz", "FizzBuzz"};
        for (int i = 1; i <= n; i++) {
            boolean t = (i % 3) == 0, f = (i % 5) == 0;
            if (!t && !f) { ans.add(String.valueOf(i)); }
            else if (t && !f) { ans.add(fb[0]); }
            else if (!t && f) { ans.add(fb[1]); }
            else { ans.add(fb[2]); }
        }
        return ans;
    }
}
