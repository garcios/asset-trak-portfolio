import {gql} from "@apollo/client";
import {Money, ValueWithPercentage} from "./get-balance-summary";

const GET_SUMMARY_TOTALS = gql`
    query GetSummaryTotals($accountId: String!){
        summaryTotals(accountId: $accountId) {
            portfolioValue{
                  amount
                  currencyCode
            }
            capitalGain{
                   amount
                   currencyCode 
                   percentage
            }
            dividends{
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
        }
    }
`;

export interface SummaryTotalsResponse {
    summaryTotals: SummaryTotalsType
}

export interface SummaryTotalsType{
    portfolioValue: Money
    capitalGain: ValueWithPercentage;
    dividends: ValueWithPercentage;
    currencyGain: ValueWithPercentage;
    totalReturn: ValueWithPercentage;
}

export default GET_SUMMARY_TOTALS;