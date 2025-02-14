import {gql} from "@apollo/client";

const GET_PERFORMANCE_NUMBERS = gql`
    query GetPerformanceNumbers($accountId: String!){
        dataItems(accountId: $accountId) {
           date
           value 
        }
    }
`;

export interface PerformanceDataResponse{
    dataItems: PerformanceData[];
}

export interface PerformanceData {
    date: string;
    value: number;
}


export default GET_PERFORMANCE_NUMBERS;