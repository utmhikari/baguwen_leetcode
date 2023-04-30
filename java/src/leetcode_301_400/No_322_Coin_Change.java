package leetcode_301_400;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.Set;

/**
 * Created by 36191 on 2018/12/1
 * 哈哈
 */
public class No_322_Coin_Change {

    public int coinChange(int[] coins, int amount) {
        if (coins == null || coins.length == 0 || amount < 1) { return 0; }
        LinkedList<Integer> queue = new LinkedList<>();
        Set<Integer> visited = new HashSet<>();
        queue.addFirst(amount);
        visited.add(amount);
        int level = 0;
        while(!queue.isEmpty()){
            int size = queue.size();
            while (size-- > 0) {
                int curr = queue.removeLast();
                if (curr == 0)  { return level; }
                if (curr > 0) {
                    for (int coin : coins) {
                        int next = curr - coin;
                        if (next >= 0 && !visited.contains(next)) {
                            queue.addFirst(next);
                            visited.add(next);
                        }
                    }
                }
            }
            level++;
        }
        return -1;
    }

    public static void main(String[] args) {
        No_322_Coin_Change solution = new No_322_Coin_Change();
        System.out.println(solution.coinChange(new int[] {186, 419, 83, 408}, 6249));
    }
}
