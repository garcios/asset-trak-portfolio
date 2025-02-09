import {GraphQLResponseResolver, HttpResponse} from "msw";

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

export default getBalanceSummaryResolver;