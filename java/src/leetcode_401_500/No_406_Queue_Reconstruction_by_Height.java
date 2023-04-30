package leetcode_401_500;

import java.util.Arrays;
import java.util.LinkedList;

/**
 * Created by 36191 on 2019/7/7
 */
public class No_406_Queue_Reconstruction_by_Height {

    public int[][] reconstructQueue(int[][] people) {
        Arrays.sort(people, (o1, o2) -> o1[0] == o2[0] ? o1[1] - o2[1] : o2[0] - o1[0]);
        LinkedList<int[]> queue = new LinkedList<>();
        for (int[] p : people) {
            queue.add(p[1], p);
        }
        return queue.toArray(new int[people.length][2]);
    }

    public static void main(String[] args) {
        No_406_Queue_Reconstruction_by_Height solution = new No_406_Queue_Reconstruction_by_Height();

    }
}
