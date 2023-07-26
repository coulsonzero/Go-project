function containsNumber1(str) {
    for(let i= 0; i< str.length; i++) {
        // !isNaN(str.charAt(i))
        // !isNaN(str[i])
        // Number(str[i])
        // item - item === 0
        if(!isNaN(Number(str[i]))) {
            return true
        }
    }
    return false
}

function containsNumber3(str) {
    for(let i = 0; i< 10; i++) {
        if(str.indexOf(i) !== -1) {
            return true
        }
    }
    return false
}

function containsNumber4(str) {
    const reg = /([0-9])/
    return reg.test(str)
}

function containsNumber(str) {
    return str.split('').some(each => !isNaN(each))
}

