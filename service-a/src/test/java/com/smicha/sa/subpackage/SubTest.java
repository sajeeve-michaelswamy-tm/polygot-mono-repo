package com.smicha.sa.subpackage;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SubTest {

    @Test
    public void someSubTest()
    {
        System.out.println("someSubTest");
        Assertions.assertEquals(1,1);
    }
}
