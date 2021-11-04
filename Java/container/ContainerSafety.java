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
 Exception in thread "main" java.lang.ClassCastException: class OrangeItem cannot be cast to class AppleItem (OrangeItem and AppleItem are in unnamed module of loader 'app')
	at ContainerSafety.main(ContainerSafety.java:21)
 */