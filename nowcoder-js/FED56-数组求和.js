

function sum(arr) {
    let res = 0

    for (let i = 0; i < arr.length; i++) {
        res += arr[i]
    }

    return res
}

function sum2(arr) {
    let res = 0
    arr.forEach(v => res += v)
    return res
}

function sum3(arr) {
    return arr.reduce((a, b) => a + b)
}

function sum4(arr) {
    return eval(arr.join("+"));
}