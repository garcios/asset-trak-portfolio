import { useQuery, gql } from "@apollo/client";
import AssetsTable from "./AssetsTable";

const GET_BALANCE_SUMMARY = gql`
    query GetBalanceSummary($accountId: String!){
        getBalanceSummary(accountId: $accountId) {
            accountId
            balanceItems {
                assetSymbol
                assetName
                price {
                    amount
                    currencyCode
                }
                quantity
                value {
                    amount
                    currencyCode
                }
                capitalGain{
                    amount
                    currencyCode
                    percentage
                }
                dividend {
                    amount
                    currencyCode
                    percentage
                }
                currencyGain{
                    amount
                    currencyCode
                    percentage
                }
                totalReturn{
                    amount
                    currencyCode
                    percentage
                }
                marketCode
            }
        }
    }
`;

interface Item {
    assetSymbol: string;
    assetName: string;
    price: number;
    priceCurrencyCode: string;
    quantity: number;
    value: number;
    valueCurrencyCode: string;
    capitalGain: number;
    capitalGainCurrencyCode: string;
    capitalGainPercentage: number;
    dividend: number;
    dividendCurrencyCode: string;
    dividendPercentage: number;
    currencyGain: number;
    currencyGainCurrencyCode: string;
    currencyGainPercentage: number;
    totalReturn: number;
    totalReturnCurrencyCode: string;
    totalReturnPercentage: number;
    marketCode: string;
}

interface Money {
    amount: number;
    currencyCode: string;
}

interface ValueWithPercentage {
    value: Money;
    percentage: number;
}


function BalanceSummary() {
    const { data, loading, error } = useQuery(GET_BALANCE_SUMMARY, {
        variables: { accountId: 'eb08df3c-958d-4ae8-b3ae-41ec04418786' }
    });

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    const items = data?.getBalanceSummary?.balanceItems.map((item: Item) => {
        return {
            assetSymbol: item.assetSymbol,
            assetName: item.assetName,
            price: item.price,
            priceCurrencyCode: item.priceCurrencyCode,
            quantity: item.quantity,
            value: item.value,
            valueCurrencyCode: item.valueCurrencyCode,
            capitalGain: item.capitalGain,
            capitalGainCurrencyCode: item.capitalGainCurrencyCode,
            capitalGainPercentage: item.capitalGainPercentage,
            dividend: item.dividend,
            dividendCurrencyCode: item.dividendCurrencyCode,
            dividendPercentage: item.dividendPercentage,
            currencyGain: item.currencyGain,
            currencyGainCurrencyCode: item.currencyGainCurrencyCode,
            currencyGainPercentage: item.currencyGainPercentage,
            totalReturn: item.totalReturn,
            totalReturnCurrencyCode: item.totalReturnCurrencyCode,
            totalReturnPercentage: item.totalReturnPercentage,
            marketCode: item.marketCode,
        }
    }).sort((a: { value: number; }, b: { value: number; }) => b.value - a.value);

    return (
        <>
            <AssetsTable items={items} />
        </>

    );
}

export default BalanceSummary;