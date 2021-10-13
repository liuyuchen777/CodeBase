/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-15 14:44:58
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-15 14:59:41
 * @Description: 
 * @FilePath: /CodeBase/Java/Basic/src/CommentStyle.java
 * @GitHub: https://github.com/liuyuchen777
 */

public class CommentStyle {
    /**
     * Java comment style:
     * multiple lines comment is like this, be careful that two * in first line
     * @see classname
     * allow user to refer other javadoc
     * @link package.class#member label
     * same as see, but use for in-line
     * @version
     * no explain, version
     * @author
     * just author
     * @docRoot
     * current label to doc root dir
     * @inheritDoc
     * @since
     * the earliest java version that the code use
     * @param parameter-name description
     * description can be multiple lines
     * @return description
     * @throws
     * @deprecated
     * HTML label is also supported by javadoc
     */

    /**
     * Java programming style:
     * [Class] The first character of class need to be upper. If it is composed by multiple words, use Camel-Case
     * [Method, Variable] the first character of word is smaller.
     */
    
    public static void main(String[] args) {
        // hello world
        System.out.println("Hello World!");
    }
}
