package alg.pattern.algorithm

typealias Validator = (String) -> Boolean

val emailValidator : Validator = { it.contains('@') }
val usernameValidator : Validator = { it.isNotEmpty() }
val passwordValidator : Validator = { it.length >= 4 }

class FormValidator(val name: String, val value: String, private val validator: Validator) {
    fun isValid() = validator(value)
}

fun Validator.optional() : Validator = { it.isEmpty() || this(it) }

fun main() {
    val emailField = FormValidator("email", "omar@test.com", emailValidator.optional())
    println("email is valid: ${emailField.isValid()}")

    val usernameField = FormValidator("username", "obarra", usernameValidator)
    println("username is valid: ${usernameField.isValid()}")

    val passwordField = FormValidator("password", "admin123", passwordValidator)
    println("password is valid: ${emailField.isValid()}")
}