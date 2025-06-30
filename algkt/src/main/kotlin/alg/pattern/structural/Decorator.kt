package alg.pattern.structural

import java.time.LocalDateTime
import java.util.UUID

fun interface Logger {
    fun log(message: String)
}

fun Logger.withUniqueId() = Logger { log("{${UUID.randomUUID()}} $it")}
fun Logger.withThreadName() = Logger { log("$it on [${Thread.currentThread().name}] thread") }
fun Logger.withDateTime(dateTime: LocalDateTime = LocalDateTime.now()) = Logger { log("[${dateTime}] - $it") }

fun main() {
    val logger = Logger { println("Console: $it") }
    logger.withUniqueId()
        .withDateTime(LocalDateTime.now())
        .withThreadName()
        .log("This is a log message")
}