package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/13
 * 043. 字符串数字相乘
 */

import java.util.Stack;

public class No_043_Multiply_Strings {

    /* See https://discuss.leetcode.com/topic/30508/easiest-java-solution-with-graph-explanation!!! */

    public String multiplyOne(String bigNum, char smallNum) {
        Stack<Character> productChars = new Stack<>();
        StringBuilder productOne = new StringBuilder();
        int smallDigit = smallNum - 48;
        int add = 0;
        int len = bigNum.length();
        for(int i = 0; i < len; i++) {
            int bigDigit = bigNum.charAt(len - 1 - i) - 48;
            int pro = bigDigit * smallDigit + add;
            add = pro / 10;
            char c = (char)(pro % 10 + 48);

            productChars.push(c);
        }
        if(add != 0) {
            productChars.push((char)(add + 48));
        }
        while(!productChars.isEmpty()) {
            productOne.append(productChars.pop());
        }
        return productOne.toString();
    }

    public String addAllProducts(String[] products) {
        Stack<Character> sumChars = new Stack<>();
        StringBuilder ans = new StringBuilder();
        int allLen = products.length;
        int index = 0;
        int add = 0;
        int[] lengths = new int[allLen];
        for(int i = 0; i < allLen; i++) {
            lengths[i] = products[i].length();
        }
        while(index < lengths[allLen - 1]) {
            int sum = add;
            for(int i = 0; i < products.length; i++) {
                if(index + 1 > lengths[i]) {
                    continue;
                }
                sum += products[i].charAt(lengths[i] - 1 - index) - 48;
            }
            add = sum / 10;
            char c = (char)(sum % 10 + 48);
            sumChars.push(c);
            index++;
        }
        if(add != 0) {
            sumChars.push((char)(add + 48));
        }
        while(!sumChars.isEmpty()) {
            ans.append(sumChars.pop());
        }
        return ans.toString();
    }

    public String multiply(String num1, String num2) {
        if(num1.equals("") || num2.equals("")) {
            return "";
        }
        if(num1.equals("0") || num2.equals("0")) {
            return "0";
        }
        int len1 = num1.length();
        int len2 = num2.length();
        int countZero = 0, countZeroNum1 = 0, countZeroNum2 = 0;
        for(int i = 0; i <= len1 - 1; i++) {
            if(num1.charAt(len1 - 1 - i) == '0') {
                countZeroNum1++;
                countZero++;
            } else {
                break;
            }
        }
        len1 -= countZeroNum1;
        num1 = num1.substring(0, len1);

        for(int i = 0; i <= len2 - 1; i++) {
            if(num2.charAt(len2 - 1 - i) == '0') {
                countZeroNum2++;
                countZero++;
            } else {
                break;
            }
        }
        len2 -= countZeroNum2;
        num2 = num2.substring(0, len2);

        String[] products = len1 < len2 ? new String[len1] : new String[len2];
        if(len1 < len2) {
            for(int i = 0; i < len1; i++) {
                StringBuilder sb = new StringBuilder();
                sb.append(multiplyOne(num2, num1.charAt(len1 - 1 - i)));
                for(int j = 0; j < i; j++) {
                    sb.append('0');
                }
                products[i] = sb.toString();
            }
        } else {
            for(int i = 0; i < len2; i++) {
                StringBuilder sb = new StringBuilder();
                sb.append(multiplyOne(num1, num2.charAt(len2 - 1 - i)));
                for(int j = 0; j < i; j++) {
                    sb.append('0');
                }
                products[i] = sb.toString();
            }
        }

        StringBuilder ans = new StringBuilder(addAllProducts(products));
        for(int i = 0; i < countZero; i++) {
            ans.append('0');
        }
        return ans.toString();
    }

    public static void main(String[] args) {
        No_043_Multiply_Strings solution = new No_043_Multiply_Strings();
        System.out.println(solution.multiply("98", "99"));
        System.out.println(solution.multiply("387632100", "38726837213210000"));
        System.out.println(solution.multiply("99999", "99999"));
    }
}
