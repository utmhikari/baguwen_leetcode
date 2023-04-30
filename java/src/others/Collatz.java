package others;

import java.math.BigInteger;

/**
 * Created by 36191 on 2018/10/5
 * Collatz 3x+1
 */
public class Collatz {

    private String val;
    private BigInteger[] ints = new BigInteger[10];

    public Collatz(String val) {
        this.val = val;
        for (int i = 0; i < ints.length; i++) {
            ints[i] = new BigInteger(String.valueOf(i));
        }
    }

    private String str(BigInteger bi) {
        StringBuilder sb = new StringBuilder();
        sb.append("Oct: ");
        sb.append(bi.toString());
        sb.append(", Bin: ");
        sb.append(bi.toString(2));
        sb.append(", Mod3: ");
        sb.append(bi.mod(ints[3]));
        sb.append(", Mod5: ");
        sb.append(bi.mod(ints[5]));
        return sb.toString();
    }

    public void output() {
        BigInteger integer = new BigInteger(val);
        while (!integer.equals(ints[1])) {
            System.out.println(str(integer));
            if (integer.mod(ints[2]).equals(ints[1])) {
                integer = integer.multiply(ints[3]).add(ints[1]);
            } else {
                integer = integer.divide(ints[2]);
            }
        }
    }

    public static void main(String[] args) {
        Collatz co = new Collatz("21");
        co.output();
    }
}
