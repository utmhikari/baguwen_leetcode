package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 057. 插入新区间
 */
import java.util.List;
import java.util.LinkedList;
public class No_057_Insert_Interval {

    public class Interval {
        int start;
        int end;
        Interval() { start = 0; end = 0; }
        Interval(int s, int e) { start = s; end = e; }
    }

    public List<Interval> insert(List<Interval> intervals, Interval newInterval) {
        intervals.add(newInterval);
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
}
