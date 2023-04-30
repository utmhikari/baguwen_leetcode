package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/29
 * 089. 二进制格雷码
 */
import java.util.List;
import java.util.ArrayList;
public class No_089_Gray_Code {

    public List<Integer> grayCode(int n) {
        List<Integer> gc = new ArrayList<>();
        for(int i = 0; i < (1 << n); i++) {
            gc.add(i ^ i >> 1);
        }
        return gc;
    }

    public void test(int n) {
        for(int i : grayCode(n)) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_089_Gray_Code solution = new No_089_Gray_Code();
        solution.test(1);
        solution.test(2);
        solution.test(3);
        solution.test(4);
        solution.test(5);
    }

}
