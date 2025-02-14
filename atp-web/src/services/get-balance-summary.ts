import {gql} from "@apollo/client";

const GET_BALANCE_SUMMARY = gql`
    query GetBalanceSummary($accountId: String!){
        BalanceSummary(accountId: $accountId) {
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

export interface BalanceSummaryResponse{
    balanceItems: InvestmentLineItem[];
}

export interface InvestmentLineItem {
    assetSymbol: string;
    assetName: string;
    price: Money;
    quantity: number;
    value: Money;
    capitalGain: MoneyWithPercentage;
    dividend: MoneyWithPercentage;
    currencyGain: MoneyWithPercentage;
    totalReturn: MoneyWithPercentage;
    marketCode: string;
}

export interface Money {
    amount: number;
    currencyCode: string;
}

export interface MoneyWithPercentage {
    amount: number;
    currencyCode: string;
    percentage: number;
}

export default GET_BALANCE_SUMMARY;