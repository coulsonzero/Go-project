/**
 * 要求在Number对象的原型对象上添加"_isPrime"函数
 * 该函数判断调用的对象是否为一个质数，是则返回true，否则返回false。
 */

Number.prototype._isPrime = function() {
    let num = this.valueOf()
    if (num < 2) return false
    for (let i = 2; i < num; i++) {
        if (num % i === 0) return false
    }
    return true
}

Number.prototype._isPrime2 = function() {
    for(let i = 2; i < this; i++) {
        if(this % i === 0) {
            return false
        }
    }
    return true
}