package leetcode_201_300;

import java.util.HashMap;
import java.util.HashSet;

/**
 * Created by 36191 on 2018/10/23
 */
public class No_290_Word_Pattern {
    public boolean wordPattern(String pattern, String str) {
        HashMap<String, Character> map = new HashMap<>(32);
        HashSet<Character> chars = new HashSet<>();
        String[] splits = str.split(" ");
        int l = pattern.length();
        if (l != splits.length) { return false; }
        for (int i = 0; i < l; i++) {
            char c = pattern.charAt(i);
            if (map.containsKey(splits[i])) { if (!map.get(splits[i]).equals(c)) { return false; }}
            else if (chars.contains(c)) { return false; }
            else { map.put(splits[i], c); chars.add(c); }
        }
        return true;
    }
}
