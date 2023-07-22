import Transaction from './Transaction';

// TODO remove after persistnet state
const demoArr = [["Saving Account","Apple",5.99],["Saving Account","Apple",5.99]];
const TransactionContainer = () => {
    return (
        <>
            {/* TODO add persistent data storage */}
            {demoArr.map((element, index) => (
                <Transaction key={index} bottom={{marginBottom:'-35px'}} account={String(element[0])} transaction={String(element[1])} amount={Number(element[2])}/>
            ))}
        </>
    );
};

export default TransactionContainer;