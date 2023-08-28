/**
 * Testing input for username creation
 *
 * @param {string} string - make sure starts with _ (if _ must have letter follow) or letter minimum and the following can be a number, letter , or underscore and limit to 25 chars
 * @returns {boolean} True or false
 */
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

/**
 * Testing input for password creation
 *
 * @param {string} string - make sure to be a  of 8 chars and can contain only 0-9A-z!@#$%^&*
 * @returns {boolean} True or false
 */
export const validatePass = (password: string) => {
    let maxChar = 8
    const complexityRegex = new RegExp(`^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*]).{${maxChar},}$`)
    if (complexityRegex.test(password)) {
        return true
    } else {
        return false
    }
}