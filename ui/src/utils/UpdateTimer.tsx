export function updateTransactionsEveryHour(lastUpdate: Date) {
    let currentDate = new Date();
    let isPastHour = currentDate.getHours() - lastUpdate.getHours()
    return isPastHour >= 1
};