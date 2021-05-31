package alg

import kotlin.test.Test
import kotlin.test.assertEquals

class WordsTest {
    @Test
    fun longestWord() {
        val got = Words.longestWord("omar barra alberto quelca")
        assertEquals("alberto", got)
    }

    @Test
    fun longestWordWithReduce() {
        val got = Words.longestWordWithReduce("omar barra alberto quelca")
        assertEquals("alberto", got)
    }
}
