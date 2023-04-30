package leetcode_201_300;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.LinkedList;

/**
 * Created by 36191 on 2018/10/23
 */
public class No_284_Peeking_Iterator {
    private class PeekingIterator implements Iterator<Integer> {

        private LinkedList<Integer> list;

        public PeekingIterator(Iterator<Integer> iterator) {
            list = new LinkedList<>();
            while (iterator.hasNext()) { list.add(iterator.next()); }
        }

        public Integer peek() {
            return list.getFirst();
        }

        @Override
        public Integer next() {
            return list.removeFirst();
        }

        @Override
        public boolean hasNext() {
            return !list.isEmpty();
        }
    }

    public PeekingIterator getPeekingIterator() {
        return new PeekingIterator(new ArrayList<Integer>().iterator());
    }
}
