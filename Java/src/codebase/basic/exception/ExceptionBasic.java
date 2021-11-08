package codebase.basic.exception;

import static codebase.basic.tool.Print.*;

class SimpleException extends Exception {}

class MyException extends Exception {
    public MyException() {}
    public MyException(String msg) {
        super(msg);
    }
}

public class ExceptionBasic {
    public void f() throws SimpleException {
        print("Throw SimpleException from f()");
        throw new SimpleException();
    }

    public static void g() throws MyException {
        print("Throw MyException from g()");
        throw new MyException();
    }

    public static void k() throws MyException {
        print("Throw MyException from k()");
        throw new MyException("Originated in k()");
    }

    public static void main(String[] args) {
        ExceptionBasic eb = new ExceptionBasic();
        try {
            eb.f();
        } catch (SimpleException e) {
            print("Caught it!");
        }

        try {
            g();
        } catch (MyException e) {
            e.printStackTrace(System.out);
        }

        try {
            k();
        } catch (MyException e) {
            e.printStackTrace(System.out);
        }
    }
}
