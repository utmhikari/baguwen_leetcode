package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 068. 文本自动校正
 */
import java.util.List;
import java.util.LinkedList;
public class No_068_Text_Justification {
    public List<String> fullJustify(String[] words, int maxWidth) {
        int maxW = maxWidth + 1;
        int sumLen = 0;
        int wordsLen = words.length;
        int tempIndex = 0;
        StringBuilder sb = new StringBuilder();
        LinkedList<String> ans = new LinkedList<>();
        for(int i = 0; i < wordsLen; i++) {
            sumLen += (words[i].length() + 1);
            if(sumLen > maxW) {
                sumLen -= (words[i].length() + 1);
                int numSpaces = maxW - sumLen;
                int numGaps = i - tempIndex - 1;
                int remainPadSpaces;
                if(numGaps != 0) {
                    int numPadSpaces = numSpaces / numGaps;
                    remainPadSpaces = numSpaces - numPadSpaces * numGaps;
                    for(int j = 0; j < numGaps; j++) {
                        sb.append(words[tempIndex + j]);
                        sb.append(' ');
                        for(int k = 0; k < numPadSpaces; k++) {
                            sb.append(' ');
                        }
                        if(remainPadSpaces > 0) {
                            sb.append(' ');
                            remainPadSpaces--;
                        }
                    }
                    sb.append(words[i - 1]);
                } else {
                    sb.append(words[i - 1]);
                    int tempLen = sb.length();
                    while(tempLen < maxWidth) {
                        sb.append(' ');
                        tempLen++;
                    }
                }
                ans.add(sb.toString());
                tempIndex = i--;
                sumLen = 0;
                sb.delete(0, sb.length());
            }
        }
        for(int i = tempIndex; i < wordsLen; i++) {
            sb.append(words[i]);
            sb.append(' ');
        }
        if(sb.length() > maxWidth) {
            sb.deleteCharAt(sb.length() - 1);
        } else {
            int tempLen = sb.length();
            while(tempLen < maxWidth) {
                sb.append(' ');
                tempLen++;
            }
        }
        ans.add(sb.toString());
        return ans;
    }

    public void test() {
        String[] words = new String[] {
                "This", "is", "an",
                "example", "of", "text",
                "justification."
        };
        List<String> ans = fullJustify(words, 16);
        for(String s : ans) {
            System.out.println(s);
        }
    }
    public static void main(String[] args) {
        No_068_Text_Justification solution = new No_068_Text_Justification();
        solution.test();
    }

}
