package utils;

import java.util.List;
import java.util.ListIterator;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Created by 36191 on 2019/1/11
 */
public class CommonUtils {
    
    public static void printStrList(List<String> l) {
        ConcurrentHashMap<String, String> map = new ConcurrentHashMap<>();
        for (String s : l) {
            System.out.print(s + ", ");
        }
        System.out.println();
    }
}
