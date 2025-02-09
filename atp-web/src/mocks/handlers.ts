import {graphql, GraphQLResponseResolver, HttpResponse} from "msw";
import getBalanceSummaryResolver from "./getBalanceSummaryResolver";

// Define the types for your mock data response
interface Money {
    amount: number;
    currencyCode: string;
}

interface BalanceItem {
    assetSymbol: string;
    assetName: string;
    price: Money;
    quantity: number;
    value: Money;
    totalGain: number;
    marketCode: string;
}

interface GetBalanceSummaryResponse {
    getBalanceSummary: {
        accountId: string;
        balanceItems: BalanceItem[];
    };
}


const handlers = [
    graphql.query<GetBalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
];

export default handlers;