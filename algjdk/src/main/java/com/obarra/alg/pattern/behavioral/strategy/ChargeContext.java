package com.obarra.alg.pattern.behavioral.strategy;

import com.obarra.alg.pattern.behavioral.strategy.algorithm.ChargeStrategy;

public class ChargeContext {
    private ChargeStrategy chargeStrategy;

    public ChargeContext(ChargeStrategy chargeStrategy) {
        this.chargeStrategy = chargeStrategy;
    }

    public void charge(final double amount) {
        chargeStrategy.charge(amount);
    }
}
