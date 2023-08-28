import { formatOverviewDate } from "../formatDate";

test('Test formatOverviewDate functionality ', () => {
    const resultOne = formatOverviewDate(new Date('2000-1-1'));
    const resultTwo = formatOverviewDate(new Date('2000-1-2'));
    const resultThree = formatOverviewDate(new Date('2000-1-3'));
    const resultFour = formatOverviewDate(new Date('2000-1-11'));
    const resultFive = formatOverviewDate(new Date('2000-1-12'));
    const resultSix = formatOverviewDate(new Date('2000-1-13'));
    const resultSeven = formatOverviewDate(new Date('2000-1-20'));
    const resultEight = formatOverviewDate(new Date('2000-1-21'));
    const resultNine = formatOverviewDate(new Date('2000-1-22'));
    const resultTen = formatOverviewDate(new Date('2000-1-23'));
    
    // test day 1
    expect(resultOne).toEqual("Jan 1st, 2000")
    // test day 2
    expect(resultTwo).toEqual("Jan 2nd, 2000")
    // test day 3
    expect(resultThree).toEqual("Jan 3rd, 2000")
    // test day 11th
    expect(resultFour).toEqual("Jan 11th, 2000")
    // test day 12th
    expect(resultFive).toEqual("Jan 12th, 2000")
    // test day 13th
    expect(resultSix).toEqual("Jan 13th, 2000")
    // test day 20th
    expect(resultSeven).toEqual("Jan 20th, 2000")
    // test day 21th
    expect(resultEight).toEqual("Jan 21st, 2000")
    // test day 22nd
    expect(resultNine).toEqual("Jan 22nd, 2000")
    // test day 23rd
    expect(resultTen).toEqual("Jan 23rd, 2000")
});