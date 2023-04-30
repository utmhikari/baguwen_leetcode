package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 071. 简化路径
 */
import java.util.LinkedList;
public class No_071_Simplify_Path {

    public String simplifyPath(String path) {
        StringBuilder sb = new StringBuilder();
        char[] pathChars = path.toCharArray();
        int startIndex = 0;
        while(startIndex < pathChars.length && pathChars[startIndex] == '/') {
            startIndex++;
        }
        int endIndex = pathChars.length - 1;
        if(startIndex == endIndex) {
            return "/";
        }
        while(endIndex >= 0 && pathChars[endIndex] == '/') {
            endIndex--;
        }
        LinkedList<String> documents = new LinkedList<>();
        for(int i = startIndex; i <= endIndex; i++) {
            if(pathChars[i] != '/' && i != endIndex) {
                sb.append(pathChars[i]);
            } else {
                if(i == endIndex) {
                    sb.append(pathChars[i]);
                }
                String doc = sb.toString();
                if(doc.equals("..")) {
                    if(!documents.isEmpty()) {
                        documents.removeLast();
                    }
                } else if(!doc.equals("") && !doc.equals(".")) {
                    documents.add(doc);
                }
                sb.delete(0, sb.length());
            }
        }
        for(String d : documents) {
            sb.append("/");
            sb.append(d);
        }
        if(sb.length() == 0) {
            return "/";
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        No_071_Simplify_Path solution = new No_071_Simplify_Path();
        System.out.println(solution.simplifyPath("/"));
        System.out.println(solution.simplifyPath("//////"));
        System.out.println(solution.simplifyPath("/home////"));
        System.out.println(solution.simplifyPath("//////home/"));
        System.out.println(solution.simplifyPath("/home/"));
        System.out.println(solution.simplifyPath("/../"));
        System.out.println(solution.simplifyPath("/a/./b/../../c/"));
        System.out.println(solution.simplifyPath("/home//foo/"));
        System.out.println(solution.simplifyPath("/./"));
    }
}
