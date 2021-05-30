package com.obarra.accurate.algorithm;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.TreeMap;

class StringCounterTest {

    @Test
    void collectWordAmounts() {
        var result = StringCounter.collectWordAmounts("java11 java11 omar barra programa en java11 barra");

        var expected = new TreeMap<String, Long>();
        expected.put("omar", 1L);
        expected.put("barra", 2L);
        expected.put("java11", 3L);
        expected.put("en", 1L);
        expected.put("programa", 1L);
        Assertions.assertIterableEquals(expected.entrySet(), result.entrySet());
    }

}