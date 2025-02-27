import {gql} from "@apollo/client";

const GET_PERFORMANCE_NUMBERS = gql`
    query GetPerformanceNumbers($accountId: String!){
        getHistoricalValues(accountId: $accountId) {
            tradeDate
            cost
            value
            currencyCode
        }
    }
`;

export interface PerformanceDataResponse{
    getHistoricalValues: PerformanceData[];
}

export interface PerformanceData {
    tradeDate: string;
    cost: number;
    value: number;
    currencyCode: string;
}


export default GET_PERFORMANCE_NUMBERS;