const _shouldUpdate1 = (oldVersion, newVersion) => {
    return newVersion > oldVersion
}

const _shouldUpdate2 = (oldVersion, newVersion) => {
    // 补全代码
    oldVersion = parseInt(oldVersion.split('.').join(''));
    newVersion = parseInt(newVersion.split('.').join(''));
    return newVersion > oldVersion;
}

const _shouldUpdate2 = (oldVersion, newVersion) => {
    oldVersion =  oldVersion.replace('.', '')
    newVersion =  newVersion.replace('.', '')
    return newVersion > oldVersion
}