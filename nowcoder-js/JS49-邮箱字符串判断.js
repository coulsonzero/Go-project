
function isAvailableEmail(sEmail) {
    const reg = /^([\w+.])+@(\w)+(\.\w+)+$/;
    return reg.test(sEmail);
}

function isAvailableEmail2(sEmail) {
    var reg = /^([\w+\.])+@(\w+)([.]\w+)+$/
    return reg.test(sEmail)
}