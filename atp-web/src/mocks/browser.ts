import {graphql, GraphQLResponseResolver, HttpResponse} from "msw";
import {setupWorker} from "msw/browser";
import {GraphQLRequest} from "@apollo/client";

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

const getBalanceSummaryResolver:GraphQLResponseResolver = ({query, variables}) => {
    console.log('Request variables:', variables);  // Log the variables

    const { accountId } = variables;
    console.log('Request accountId:', accountId);

    const mockData = { data: {
            getBalanceSummary: {
                accountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
                balanceItems: [
                    {
                        assetSymbol: "AAPL",
                        assetName: "Apple Inc.",
                        price: {amount: 150.0, currencyCode: "USD"},
                        quantity: 10,
                        value: {amount: 1500.0, currencyCode: "USD"},
                        totalGain: 50,
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "GOOGL",
                        assetName: "Alphabet Inc.",
                        price: {amount: 2800.0, currencyCode: "USD"},
                        quantity: 2,
                        value: {amount: 5600.0, currencyCode: "USD"},
                        totalGain: 1500,
                        marketCode: "NASDAQ",
                    },
                ],
            },
        }};

    // Return mock data
    return HttpResponse.json(mockData)
};

const handlers = [
    graphql.query<GetBalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
];


// Create the worker and define the GraphQL query handler
const worker = setupWorker(...handlers);

export default worker;