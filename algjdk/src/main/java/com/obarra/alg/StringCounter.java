package com.obarra.alg;

import java.util.Map;
import java.util.TreeMap;

public class StringCounter {

    public static Map<String, Long> collectWordAmounts(final String phrase) {
        var wordCollector = new TreeMap<String, Long>();
        var words = phrase.split(" ");

        for (String word : words) {
            wordCollector.compute(word, (key, value)
                    -> (value == null) ? 1 : value + 1);
        }
        return wordCollector;
    }
}
