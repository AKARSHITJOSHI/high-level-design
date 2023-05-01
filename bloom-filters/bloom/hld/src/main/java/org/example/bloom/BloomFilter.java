package org.example.bloom;

import java.math.BigInteger;
import java.nio.ByteBuffer;
import java.security.*;
public class BloomFilter implements Filter{

    //set is our structure where we set the bits
    private byte[] filter;

    //keySize refers to each object size ,filter is the size of bloom filter , and size is no.OfObjects
    private int keySize, filterSize, objectCount;
    private MessageDigest md;
    public BloomFilter(int capacity, int k)
    {
        filterSize = capacity;
        filter = new byte[filterSize];
        keySize = k;
        objectCount = 0;
        try
        {
            md =MessageDigest.getInstance(Algorithm.MD5.name);
        }
        catch (NoSuchAlgorithmException e)
        {
            throw new IllegalArgumentException("Error : MD5 Hash not found");
        }
    }

    public void clear()
    {
        filter = new byte[filterSize];
        objectCount = 0;
        try
        {
            md = MessageDigest.getInstance(Algorithm.MD5.name);
        }
        catch (NoSuchAlgorithmException e)
        {
            throw new IllegalArgumentException("Error : MD5 Hash not found");
        }
    }

    public boolean isEmpty()
    {
        return objectCount == 0;
    }

    public int getObjectCount()
    {
        return objectCount;
    }

    //get the array and map each bit to 1
    @Override
    public void store(Object obj) throws Exception {
        int[] tmpset = getSetArray(obj);
        for (int i : tmpset)
            filter[i] = 1;
        objectCount++;
    }

    //if any of the bit not ON within the tempset we return false
    @Override
    public boolean isPresent(Object obj) {
        int[] tmpset = getSetArray(obj);
        for (int i : tmpset)
            if (filter[i] != 1)
                return false;
        return true;
    }
    private int[] getSetArray(Object obj)
    {
        int[] tmpset = new int[keySize];
        tmpset[0] = getHash(obj.hashCode());
        for (int i = 1; i < keySize; i++)
            //sets hash based on previous index
            tmpset[i] = (getHash(tmpset[i - 1]));
        return tmpset;
    }

    //returns an integer hash based on hashcode passed .
    private int getHash(int i)
    {
        md.reset();
        byte[] bytes = ByteBuffer.allocate(4).putInt(i).array();
        md.update(bytes, 0, bytes.length);
        return Math.abs(new BigInteger(1, md.digest()).intValue()) % (filter.length - 1);
    }

    //Limitations cannot delete elements .
}
