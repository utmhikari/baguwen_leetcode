package leetcode_001_100;
/**
 * Created by 36191 on 2017/10/5.
 * 013.罗马数字到整数转换
 */

public class No_013_Roman_To_Integer {

    public int romanToInt(String s) {
        int value = 0;
        int index = 0;
        int len = s.length();
        String[] ones = new String[]{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
        String[] tens = new String[]{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
        String[] hundreds = new String[]{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
        String[] thousands = new String[]{"M", "MM", "MMM"};
        int onesLength = ones.length;
        int tensLength = tens.length;
        int hundredsLength = hundreds.length;
        int thousandsLength = thousands.length;
        int digitlen;
        for(int i = thousandsLength - 1; i >= 0; i--) {
            digitlen = thousands[i].length();
            if(index + digitlen > len) {
                continue;
            }
            if(thousands[i].equals(s.substring(index, index + digitlen))) {
                value += 1000 * (i + 1);
                index += digitlen;
                if(index == len) {
                    return value;
                }
                break;
            }
        }
        for(int i = hundredsLength - 1; i >= 0; i--) {
            digitlen = hundreds[i].length();
            if(index + digitlen > len) {
                continue;
            }
            if(hundreds[i].equals(s.substring(index, index + digitlen))) {
                value += 100 * (i + 1);
                index += digitlen;
                if(index == len) {
                    return value;
                }
                break;
            }
        }
        for(int i = tensLength - 1; i >= 0; i--) {
            digitlen = tens[i].length();
            if(index + digitlen > len) {
                continue;
            }
            if(tens[i].equals(s.substring(index, index + digitlen))) {
                value += 10 * (i + 1);
                index += digitlen;
                if(index == len) {
                    return value;
                }
                break;
            }
        }
        for(int i = onesLength - 1; i >= 0; i--) {
            digitlen = ones[i].length();
            if(index + digitlen > len) {
                continue;
            }
            if(ones[i].equals(s.substring(index, index + digitlen))) {
                value += (i + 1);
                index += digitlen;
                if(index == len) {
                    return value;
                }
                break;
            }
        }
        return value;
    }

    public void output(String s) {
        System.out.println(romanToInt(s));
    }

    public static void main(String[] args) {
        No_013_Roman_To_Integer solution = new No_013_Roman_To_Integer();
        solution.output("I");
        solution.output("IV");
        solution.output("X");
        solution.output("XIX");
        solution.output("XX");
        solution.output("XLV");
        solution.output("C");
        solution.output("CDXCIX");
        solution.output("M");
        solution.output("MCMLXXX");
        solution.output("MMCMXCIX");
    }
}
