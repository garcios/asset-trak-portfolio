import {gql} from "@apollo/client";

const GET_BALANCE_SUMMARY = gql`
    query GetBalanceSummary($accountId: String!){
        BalanceSummary(accountId: $accountId) {
            BalanceItems {
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

export interface BalanceSummaryResponse{
    balanceItems: LineItem[];
};

export interface LineItem {
    assetSymbol: string;
    assetName: string;
    price: Money;
    quantity: number;
    value: Money;
    capitalGain: ValueWithPercentage;
    dividend: ValueWithPercentage;
    currencyGain: ValueWithPercentage;
    totalReturn: ValueWithPercentage;
    marketCode: string;
};

export interface Money {
    amount: number;
    currencyCode: string;
};

export interface ValueWithPercentage {
    amount: number;
    currencyCode: string;
    percentage: number;
};

export default GET_BALANCE_SUMMARY;