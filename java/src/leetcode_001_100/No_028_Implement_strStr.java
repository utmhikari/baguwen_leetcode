package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/9.
 * 028.实现寻找子串
 */
public class No_028_Implement_strStr {

    /*
    public int strStr(String haystack, String needle) {
        for (int i = 0; ; i++) {
            for (int j = 0; ; j++) {
                if (j == needle.length()) return i;
                if (i + j == haystack.length()) return -1;
                if (needle.charAt(j) != haystack.charAt(i + j)) break;
            }
        }
    }
     */

    public int strStr(String haystack, String needle) {
        if(needle.equals("")) {
            return 0;
        }
        int index = -1;
        int nextTempIndex = 0;
        int hsLength = haystack.length();
        int nLength = needle.length();
        char firstLetter = needle.charAt(0);
        int i = 0;
        while(i < hsLength - nLength + 1) {
            if(haystack.charAt(i) != firstLetter) {
                i++;
            } else {
                boolean flag = true;
                boolean next = false;
                for(int j = 0; j < nLength; j++) {
                    char h = haystack.charAt(i + j);
                    char n = needle.charAt(j);
                    if(!next && j > 0 && h == firstLetter) {
                        next = true;
                        nextTempIndex = i + j;
                    }
                    if(flag && h != n) {
                        flag = false;
                    }
                    if(!flag && next) {
                        break;
                    }
                }
                if(!flag) {
                    if(next) {
                        i = nextTempIndex;
                    } else {
                        i += nLength;
                    }
                } else {
                    return i;
                }
            }
        }
        return index;
    }

    public static void main(String[] args) {
        No_028_Implement_strStr solution = new No_028_Implement_strStr();
        System.out.println(solution.strStr("1", "2"));
        System.out.println(solution.strStr("123", "23"));
        System.out.println(solution.strStr("123", "123"));
        System.out.println(solution.strStr("123", "1234"));
        System.out.println(solution.strStr("12323", "23"));
        System.out.println(solution.strStr("mississippi", "issip"));
        System.out.println(solution.strStr("aabaaabaaac", "aabaaac"));
    }
}
