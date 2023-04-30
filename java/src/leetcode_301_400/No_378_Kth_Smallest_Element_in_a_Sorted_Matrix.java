package leetcode_301_400;

import java.util.Comparator;
import java.util.HashSet;
import java.util.PriorityQueue;

/**
 * Created by 36191 on 2019/2/8
 */
public class No_378_Kth_Smallest_Element_in_a_Sorted_Matrix {

    private int getPos(int x, int y, int w, int h) {
        if (x >= w || y >= h) { return -1; }
        return w * y + x;
    }

    private int[] getXY(int pos, int w) {
        return new int[] { pos % w, pos / w };
    }

    public int kthSmallest(int[][] matrix, int k) {
        HashSet<Integer> visited = new HashSet<>();
        int h = matrix.length, w = matrix[0].length;
        PriorityQueue<Integer[]> ans = new PriorityQueue<>(
                Comparator.comparingInt((Integer[] xy) -> matrix[xy[1]][xy[0]]).reversed()
        );
        PriorityQueue<Integer[]> subs = new PriorityQueue<>(
                Comparator.comparingInt((Integer[] xy) -> matrix[xy[1]][xy[0]])
        );
        int startPos = getPos(0, 0, w, h);
        visited.add(startPos);
        subs.offer(new Integer[] { 0, 0 });
        while (true) {
            if (subs.isEmpty()) { break; }
            Integer[] peek = subs.peek();
            if (ans.size() < k) {
                ans.offer(peek);
                subs.poll();
            } else {
                Integer[] ansXY = ans.peek();
                if (matrix[ansXY[1]][ansXY[0]] <= matrix[peek[1]][peek[0]]) { break; }
                ans.poll();
                ans.offer(peek);
                subs.poll();
            }
            int posRight = getPos(peek[0] + 1, peek[1], w, h);
            if (posRight >= 0 && !visited.contains(posRight)) {
                visited.add(posRight);
                subs.offer(new Integer[] { peek[0] + 1, peek[1] });
            }
            int posDown = getPos(peek[0], peek[1] + 1, w, h);
            if (posDown >= 0 && !visited.contains(posDown)) {
                visited.add(posDown);
                subs.offer(new Integer[] { peek[0], peek[1] + 1 });
            }
        }
        Integer[] xy = ans.poll();
        return matrix[xy[1]][xy[0]];
    }

    public static void main(String args) {

    }
}
