package org.example.bloom;

public interface Filter {

    public  boolean isPresent(Object obj);
    public <T> void store(Object obj) throws Exception;
}
