import { useQuery, gql } from "@apollo/client";
import AssetsTable from "./AssetsTable";

const GET_BALANCE_SUMMARY = gql`
    query {
        getBalanceSummary(accountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786") {
            accountId
            totalValue {
                amount
                currencyCode
            }
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
                totalGain
                marketCode
            }
        }
    }
`;

interface Item {
    assetName: string;
    assetSymbol: string;
    price: Money;
    quantity: number;
    value: Money;
    totalGain: number;
    marketCode: string;
}

interface Money {
    amount: number;
    currencyCode: string;
}

function BalanceSummary() {
    const {loading, error, data} = useQuery(GET_BALANCE_SUMMARY);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    const items = data?.getBalanceSummary?.balanceItems.map((item: Item) => {
        return {
            assetSymbol: item.assetSymbol,
            assetName: item.assetName,
            price: item.price.amount,
            currency: item.price.currencyCode,
            quantity: item.quantity,
            value: item.value.amount,
            totalGain: 0,
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