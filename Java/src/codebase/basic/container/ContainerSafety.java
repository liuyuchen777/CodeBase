package codebase.basic.container;

import java.util.*;

class AppleItem {
    private static long counter;
    private final long id = counter++;
    public long id() { return id; }
}

class OrangeItem {

}

public class ContainerSafety {
    @SuppressWarnings("unchecked")
    public static void main(String[] args) {
        ArrayList apples = new ArrayList();
        for (int i = 0; i < 3; i++)
            apples.add(new AppleItem());
        apples.add(new OrangeItem());
        for (int i = 0; i < apples.size(); i++)
            ((AppleItem)apples.get(i)).id();
    }
}

/*
 Exception in thread "main" java.lang.ClassCastException: class codebase.basic.container.OrangeItem cannot be cast to class codebase.basic.container.AppleItem (codebase.basic.container.OrangeItem and codebase.basic.container.AppleItem are in unnamed module of loader 'app')
	at codebase.basic.container.ContainerSafety.main(codebase.basic.container.ContainerSafety.java:21)
 */