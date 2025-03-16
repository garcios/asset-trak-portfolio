import {gql} from "@apollo/client";

const GET_PERFORMANCE_NUMBERS = gql`
    query GetPerformanceNumbers($accountId: String!, $startDate: String!, $endDate: String!){
        performanceHistory(accountId: $accountId, startDate: $startDate, endDate: $endDate) {
            tradeDate
            cost
            value
            currencyCode
        }
    }
`;

export interface PerformanceDataResponse{
    performanceHistory: PerformanceData[];
}

export interface PerformanceData {
    tradeDate: string;
    cost: number;
    value: number;
    currencyCode: string;
}


export default GET_PERFORMANCE_NUMBERS;