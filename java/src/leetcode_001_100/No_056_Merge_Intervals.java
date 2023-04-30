package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 056. 合并区间
 */
import java.util.List;
import java.util.LinkedList;
import java.util.ArrayList;
public class No_056_Merge_Intervals {

    public class Interval {
       int start;
       int end;
       Interval() { start = 0; end = 0; }
       Interval(int s, int e) { start = s; end = e; }
    }

    public List<Interval> merge(List<Interval> intervals) {
        int size = intervals.size();
        if(size <= 1) {
            return intervals;
        }
        intervals.sort((i1, i2) -> Integer.compare(i1.start, i2.start));
        List<Interval> merge = new LinkedList<>();
        int start = intervals.get(0).start;
        int end = intervals.get(0).end;
        for(Interval itv : intervals) {
            if (itv.start <= end)
                end = Math.max(end, itv.end);
            else {
                merge.add(new Interval(start, end));
                start = itv.start;
                end = itv.end;
            }
        }
        merge.add(new Interval(start, end));
        return merge;
    }

    /*
    public List<Interval> merge(List<Interval> intervals) {
        int size = intervals.size();
        if(size <= 1) {
            return intervals;
        }
        boolean canChange = false;
        for(int i = 0; i < size - 1; i++) {
            for(int j = i + 1; j < size; j++) {
                int iStart = intervals.get(i).start;
                int iEnd = intervals.get(i).end;
                int jStart = intervals.get(j).start;
                int jEnd = intervals.get(j).end;
                if((iStart <= jEnd && iEnd >= jEnd) || (iEnd >= jStart && iStart <= jStart)) {
                    canChange = true;
                    intervals.remove(i);
                    intervals.remove(j - 1);
                    intervals.add(new Interval(Math.min(iStart, jStart), Math.max(iEnd, jEnd)));
                    break;
                }
            }
            if(canChange) {
                break;
            }
        }
        if(!canChange) {
            return intervals;
        }
        return merge(intervals);
    }
    */

    public void test() {
        List<Interval> intervals = new ArrayList<Interval>() {{
            add(new Interval(1, 3));
            add(new Interval(15, 18));
            add(new Interval(2, 6));
            add(new Interval(8, 10));
        }};
        List<Interval> merge = merge(intervals);
        for(Interval itv : merge) {
            System.out.print(itv.start);
            System.out.print(", ");
            System.out.print(itv.end);
            System.out.print("\n");
        }
    }

    public static void main(String[] args) {
        No_056_Merge_Intervals solution = new No_056_Merge_Intervals();
        solution.test();
    }
}
