import { validateUser, validatePass } from "./validator";

test('Test validator functionality ', () => {
    const resultOne = validatePass("P@ssw0rd");
    const resultTwo = validatePass("Password");
    const resultThree = validateUser("user");
    const resultFour = validateUser("1user");
    
    // test if password pass validation
    expect(resultOne).toEqual(true)
    // test if password fails validation
    expect(resultTwo).toEqual(false)
    // test if username pass validation
    expect(resultThree).toEqual(true)
    // test if username fails validation
    expect(resultFour).toEqual(false)
});