package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 060. 排列组合顺序
 */
import java.util.LinkedList;
public class No_060_Permutation_Sequence {

    public String getPermutation(int n, int k) {
        StringBuilder sb = new StringBuilder();
        int[] facts = new int[n + 1];
        int fact = 1;
        facts[0] = fact;
        LinkedList<Integer> nums = new LinkedList<>();
        for(int i = 1; i <= n; i++) {
            fact *= i;
            facts[i] = fact;
            nums.add(i);
        }
        k--;
        for(int i = 1; i <= n; i++) {
            int index = k / facts[n - i];
            sb.append(nums.get(index));
            nums.remove(index);
            k -= index * facts[n - i];
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        No_060_Permutation_Sequence solution = new No_060_Permutation_Sequence();
        System.out.println(solution.getPermutation(3, 4));
    }
}
