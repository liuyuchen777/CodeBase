import java.lang.annotation.*;

/*
 * @Author: Liu Yuchen
 * @Date: 2021-06-28 07:58:29
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-06-28 08:04:07
 * @Description: 
 * @FilePath: /CodeBase/Java/Annotation/src/InheritableAnnotation.java
 * @GitHub: https://github.com/liuyuchen777
 */

/**
 * 这个自定注解表示类是否能够继承
 * @Documented
 * @Retention(RetentionPolicy.RUNTIME)
 * @Target(ElementType.ANNOTATION_TYPE)
 * public @interface Inherited {
 * }
 */

@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Inherited // 如果注释掉，子类为false
@interface Inheritable
{
}

@Inheritable
class InheritableFather
{
    public InheritableFather() {
        // InheritableBase是否具有 Inheritable Annotation
        System.out.println("InheritableFather:" + 
            InheritableFather.class.isAnnotationPresent(Inheritable.class));
    }
}

class InheritableSon extends InheritableFather
{
    public InheritableSon() {
        super();    // 调用父类的构造函数
        // InheritableSon类是否具有 Inheritable Annotation
        System.out.println("InheritableSon:" + 
            InheritableSon.class.isAnnotationPresent(Inheritable.class));
    }
}

public class InheritableAnnotation {
    public static void main(String[] args) {
        InheritableSon is = new InheritableSon();
    }
}
