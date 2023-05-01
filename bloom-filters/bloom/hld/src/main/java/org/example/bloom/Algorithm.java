package org.example.bloom;

public enum Algorithm {

    MD5("MD5"),
    MD2("MD2"),
    SHA1("SHA-1");
    String name;

     Algorithm(String n){
        this.name=n;
    }
}
