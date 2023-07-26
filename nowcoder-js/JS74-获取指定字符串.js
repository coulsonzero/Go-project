


function captureThreeNumbers1(str) {
    let arr = str.match(/\d{3}/)
    return arr ? arr[0] : false
}

function captureThreeNumbers2(str) {
    let reg = /\d{3}/
    if (!reg.test(str)) return false
    return str.match(reg)[0]
}

function captureThreeNumbers3(str) {
    return /\d{3}/.exec(str) || false
}