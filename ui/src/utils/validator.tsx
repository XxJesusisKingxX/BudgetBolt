export const validateUser = (username: string) => {
    let maxChar = 25
    const validStart = new RegExp(`^_?[a-zA-Z][a-zA-Z0-9_]{1,${maxChar}}$`); // make sure starts with _ (if _ must have letter follow) or letter minimum and the following can be a number, letter , or underscore
    const isUnder = username.length <= maxChar ? true : false;
    if (validStart.test(username) && isUnder) {
        return true
    } else {
        return false
    }
}

export const validatePass = (password: string) => {
    let maxChar = 8
    const complexityRegex = new RegExp(`^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*]).{${maxChar},}$`)
    if (complexityRegex.test(password)) {
        return true
    } else {
        return false
    }
}