package com.obarra.alg.pattern.behavioral.strategy;

import java.util.ArrayList;
import java.util.List;

import com.obarra.alg.pattern.behavioral.strategy.algorithm.ChargeStrategy;
import com.obarra.alg.pattern.behavioral.strategy.algorithm.DebitAccountStrategy;
import com.obarra.alg.pattern.behavioral.strategy.algorithm.MercadoPagoStrategy;

public class DemoStrategy {
    public static void main(String[] args) {
        ChargeContext chargeContext = new ChargeContext(new MercadoPagoStrategy());
        chargeContext.charge(404);

        chargeContext = new ChargeContext(new DebitAccountStrategy());
        chargeContext.charge(501);

        List<ChargeStrategy> chargeStrategies = getMethodOfCharge();
        chargeStrategies.forEach(context -> {
            new ChargeContext(context).charge(23);
        });


    }

    private static List<ChargeStrategy> getMethodOfCharge() {
        List<ChargeStrategy> chargeStrategies = new ArrayList<>();
        chargeStrategies.add(System.out::println);
        chargeStrategies.add(x -> {
            System.out.println(x * 2);
        });
        chargeStrategies.add(x -> {
            System.out.println(x * 5);
        });

        return chargeStrategies;
    }
}
