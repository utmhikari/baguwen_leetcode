package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/15
 * 049.组合易位构词
 */
import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;
public class No_049_Group_Anagrams {

    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> groups = new ArrayList<>();
        int len = strs.length;
        if(len == 0) {
            return groups;
        }
        HashMap<Integer, List<String>> groupsMap = new HashMap<>();
        for(String s : strs) {
            int sLen = s.length();
            int hash = sLen;
            for(int i = 0; i < sLen; i++) {
                int c = s.charAt(i);
                hash += c * (c + 7) * (c + 31) * (c + 127);
            }
            if(groupsMap.isEmpty() || !groupsMap.containsKey(hash)) {
                groupsMap.put(hash, new ArrayList<String>(){{ add(s); }});
            } else {
                groupsMap.get(hash).add(s);
            }
        }
        for(int h : groupsMap.keySet()) {
            groups.add(groupsMap.get(h));
        }
        return groups;
    }

    public void test(String[] strs) {
        for(List<String> l : groupAnagrams(strs)) {
            for(String s : l) {
                System.out.print(s + ", ");
            }
            System.out.println();
        }
    }
    public static void main(String[] args) {
        No_049_Group_Anagrams solution = new No_049_Group_Anagrams();
        solution.test(new String[] {"eat", "tea", "tan", "ate", "nat", "bat"});
    }
}
