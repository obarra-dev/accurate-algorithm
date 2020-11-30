package com.obarra.accurate.algorithm;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class EvenOddTest {

    @Test
    void isEven() {
        assertTrue(EvenOdd.isEven(0L));
    }

    @Test
    void isOdd() {
        assertFalse(EvenOdd.isOdd(0L));
    }

    @Test
    void isEvenWhenUseMaxLong() {
        assertFalse(EvenOdd.isEven(Long.MAX_VALUE));
    }

    @Test
    void isEvenStrategyBitwiseANDOperator() {
        assertTrue(EvenOdd.isEvenStrategyBitwiseANDOperator(0L));
    }

    @Test
    void isOddStrategyBitwiseANDOperator() {
        assertFalse(EvenOdd.isOddStrategyBitwiseANDOperator(0L));
    }

    @Test
    void isEvenStrategyBitwiseANDOperatorWhenUseMaxLong() {
        assertFalse(EvenOdd.isEvenStrategyBitwiseANDOperator(Long.MAX_VALUE));
    }


    @Test
    void isEvenStrategyDivisionOperator() {
        assertTrue(EvenOdd.isEvenStrategyDivisionOperator(0L));
    }

    @Test
    void isOddStrategyDivisionOperator() {
        assertFalse(EvenOdd.isOddStrategyDivisionOperator(0L));
    }

    @Test
    void isEvenStrategyDivisionOperatorWhenUseMaxLong() {
        assertFalse(EvenOdd.isEvenStrategyDivisionOperator(Long.MAX_VALUE));
    }

}