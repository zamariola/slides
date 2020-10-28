typealias operation = (Double) -> Double 

fun sumWith(toSum: Double): operation {
    return {x: Double -> x + toSum}
}

fun subtract(toSub: Double): operation {
    return {x: Double -> x - toSub}
}

fun divideBy(denominator: Double): operation {
    return {x: Double -> x / denominator}
}

fun main(args: Array<String>) {
 
    val result = calculateFunc(arrayOf(
        sumWith(3.0),
        sumWith(4.5),
        subtract(2.1),
        divideBy(2.0),
        ))
    println(result)
}

fun calculateFunc(ops: Array<operation>): Double {
    return ops.fold(0.0) {result, op -> op(result)}
}

fun calculate(ops: Array<operation>): Double {
    var memory = 0.0
    for (op in ops) {
        memory = op(memory)
    }
    return memory
}

