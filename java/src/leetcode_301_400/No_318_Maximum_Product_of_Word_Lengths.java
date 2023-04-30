package leetcode_301_400;

/**
 * Created by 36191 on 2018/12/1
 */
public class No_318_Maximum_Product_of_Word_Lengths {

    public int maxProduct(String[] words) {
        int max = 0, len = words.length;
        if (len <= 1) { return max; }
        int[] letters = new int[len], lens = new int[len];
        for (int i = 0; i < len; i++) {
            lens[i] = words[i].length();
            for (int j = 0; j < lens[i]; j++) {
                int idx = words[i].charAt(j) - 'a';
                if (((1 << idx) | letters[i]) != letters[i]) { letters[i] |= (1 << idx); }
            }
        }
        for(int i = 0; i < len - 1; i++) {
            for (int j = i + 1; j < len; j++) {
                if ((letters[i] & letters[j]) == 0) {
                    max = Math.max(max, lens[i] * lens[j]);
                }
            }
        }
        return max;
    }

    public static void main(String[] args) {
        No_318_Maximum_Product_of_Word_Lengths solution = new No_318_Maximum_Product_of_Word_Lengths();
        String[][] inputs = new String[][] {
                new String[] {"abcw","baz","foo","bar","xtfn","abcdef"},
                new String[] {"a","ab","abc","d","cd","bcd","abcd"},
                new String[] {"a","aa","aaa","aaaa"},
        };
        for (String[] input : inputs) {
            System.out.println(solution.maxProduct(input));
        }
    }
}
