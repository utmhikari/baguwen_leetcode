package leetcode_201_300;

import java.util.LinkedList;

/**
 * Created by 36191 on 2018/10/6
 */
public class No_215_Kth_Largest_Element_in_an_Array {
    
    private void insert(int num, LinkedList<Integer> list) {
        if (num < list.getFirst()) { list.addFirst(num); }
        else if (num >= list.getLast()) { list.addLast(num); }
        else {
            for (int j = 0; j < list.size(); j++) {
                if (num <= list.get(j)) { list.add(j, num); break; }
            }
        }
    }
    public int findKthLargest(int[] nums, int k) {
        LinkedList<Integer> list = new LinkedList<Integer>() {{ add(nums[0]); }};
        // sort k elements at front
        for (int i = 1; i < k; i++) { insert(nums[i], list); }
        // insert remained n - k elements
        for (int i = k; i < nums.length; i++) {
            if (nums[i] > list.getFirst()) { insert(nums[i], list); list.removeFirst(); }
        }
        return list.getFirst();
    }
}
