package codebase.basic.io;/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-25 13:30:58
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-25 13:56:13
 * @Description: 
 * @FilePath: /CodeBase/Java/IO/src/codebase.basic.io.BasicIO.java
 * @GitHub: https://github.com/liuyuchen777
 */
import java.util.regex.*;
import java.io.*;
import java.util.*;

class DirFilter implements FilenameFilter {
    private Pattern pattern;
    public DirFilter(String regex) {
        pattern = Pattern.compile(regex);
    }
    public boolean accept(File dir, String name) {
        return pattern.matcher(name).matches();
    }
}

public class BasicIO {
    public static void main(String[] args) {
        File path = new File(".");
        String[] list;
        if (args.length == 0)
            list = path.list();
        else 
            list = path.list(new DirFilter(args[0]));
        Arrays.sort(list, String.CASE_INSENSITIVE_ORDER);
        for (String dirItem : list)
            System.out.println(dirItem);
    }
}
