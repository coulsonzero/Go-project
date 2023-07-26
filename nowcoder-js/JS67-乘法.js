/**
 * 求 a 和 b 相乘的值，a 和 b 可能是小数，需要注意结果的精度问题
 * input: 3, 0.0001
 * output: 0.0003
 */

function multiply(a, b) {
    return Number((a * b).toFixed(4))
}


console.log(multiply(3, 0.1))       // 0.3
console.log(multiply(3, 0.0001))    // 0.0003