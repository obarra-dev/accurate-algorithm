package com.obarra.alg;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class FibonacciTest {

    private Integer seed;

    private List<Integer> expected;

    @BeforeEach
    public void setUp() {
        seed = 10;
        expected = Arrays.asList(0, 1, 1, 2, 3, 5, 8, 13, 21, 34);
    }

    @Test
    public void generateByStream() {
        assertEquals(expected, Fibonacci.generate(10));
    }

    @Test
    public void generate() {
        assertEquals(expected, Fibonacci.generateByStream(10));
    }
}