package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/26
 */
import java.util.*;

public class No_134_Gas_Station {

    public int canCompleteCircuit(int[] gas, int[] cost) {
        if (gas.length == 0) { return -1; }
        int sumSubtract = 0;
        int minusIdx = -1;
        HashSet<Integer> starts = new HashSet<>();
        int[] subtract = new int[gas.length];
        for (int i = 0; i < gas.length; i++) {
            subtract[i] = gas[i] - cost[i];
            sumSubtract += subtract[i];
            if (subtract[i] > 0 && minusIdx != -1) {
                starts.add(i);
                minusIdx = -1;
            } else {
                minusIdx = i;
            }
        }
        if (sumSubtract < 0) { return -1; }
        if (subtract[0] > 0 && subtract[subtract.length - 1] <= 0 || starts.size() == 0) {
            starts.add(0);
        }
        for (int s : starts) {
            int sum = 0;
            boolean flag = true;
            for (int i = 0; i < subtract.length; i++) {
                int idx = (s + i) % subtract.length;
                sum += subtract[idx];
                if(sum < 0) {
                    flag = false;
                    break;
                }
            }
            if (flag) {
                return s;
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        No_134_Gas_Station solution = new No_134_Gas_Station();
        int[] gas1 = new int[] {1,2,3,4,5};
        int[] cost1 = new int[] {3,4,5,1,2};
        int[] gas2 = new int[] {2,3,4};
        int[] cost2 = new int[] {3,4,3};
        int[] gas3 = new int[] {3};
        int[] cost3 = new int[] {2};
        System.out.println(solution.canCompleteCircuit(gas1, cost1));
        System.out.println(solution.canCompleteCircuit(gas2, cost2));
        System.out.println(solution.canCompleteCircuit(gas3, cost3));
    }
}
