package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_365_Water_and_Jug_Problem {

    public int getGCD(int x, int y) {
        if (x == 0) { return y; }
        return getGCD(y % x, x);
    }
    
    public boolean calculateGCD(int x, int y, int k) {
        return k % getGCD(x,y) == 0;
    }
    
    public boolean canMeasureWater(int x, int y, int z) {
        if (x + y < z) { return false; }
        if (x + y == z) { return true; }
        return calculateGCD(x, y, z);
    }
}
