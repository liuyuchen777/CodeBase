package com.example.jcip;

import net.jcip.annotations.*;

@NotThreadSafe
public class UnsafeSequence {
    private int value;
    public int getNext() {
        return value++;
    }
}
