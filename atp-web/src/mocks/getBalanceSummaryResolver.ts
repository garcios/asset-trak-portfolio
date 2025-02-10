import {GraphQLResponseResolver, HttpResponse} from "msw";

// Define the types for your mock data response
interface Money {
    amount: number;
    currencyCode: string;
}

interface ValueWithPercentage {
    value: Money;
    percentage: number;
}

interface BalanceItem {
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
                        assetSymbol: "IVV",
                        assetName: "iShares S&P 500 ETF Trust",
                        price: {amount: 64.21, currencyCode: "AUD"},
                        quantity: 1442,
                        value: {amount: 92590.82, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 18935.30, currencyCode: "AUD"},
                            percentage: 25.71,
                        },
                        dividend:{
                            value: { amount: 862.58, currencyCode: "AUD"},
                            percentage: 1.17,
                        },
                        currencyGain: {
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        totalReturn: {
                            value: { amount: 19797.88, currencyCode: "AUD"},
                            percentage: 26.88,
                        },
                        marketCode: "ASX",
                    },
                    {
                        assetSymbol: "AMZN",
                        assetName: "Amazon.com Inc.",
                        price: {amount: 229.15, currencyCode: "USD"},
                        quantity: 116,
                        value: {amount: 42508.37, currencyCode: "AUD"},
                        capitalGain: {
                             value: { amount: 10957.74, currencyCode: "AUD"},
                             percentage: 36.78,
                         },
                        dividend:{
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        currencyGain:{
                            value: { amount: 1711.50, currencyCode: "AUD"},
                            percentage: 5.75,
                        },
                        totalReturn: {
                            value: { amount: 12669.24, currencyCode: "AUD"},
                            percentage: 42.53,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "GOOGL",
                        assetName: "Alphabet Inc - Class A",
                        price: {amount: 185.34, currencyCode: "USD"},
                        quantity: 115,
                        value: {amount: 34068.23, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 4393.17, currencyCode: "AUD"},
                            percentage: 15.72,
                        },
                        dividend: {
                            value: { amount: 69.28, currencyCode: "AUD"},
                            percentage: 0.25,
                        },
                        currencyGain: {
                            value: { amount: 1703.98, currencyCode: "AUD"},
                            percentage: 6.10,
                        },
                        totalReturn: {
                            value: { amount: 6166.43, currencyCode: "AUD"},
                            percentage: 22.06,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "MSFT",
                        assetName: "Microsoft Inc.",
                        price: {amount: 409.75, currencyCode: "USD"},
                        quantity: 49,
                        value: {amount: 32092.06, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 282.99, currencyCode: "AUD"},
                            percentage: 0.94,
                        },
                        dividend: {
                            value: { amount: 139.37, currencyCode: "AUD"},
                            percentage: 0.46,
                        },
                        currencyGain: {
                            value: { amount: 1675.89, currencyCode: "AUD"},
                            percentage: 5.57,
                        },
                        totalReturn: {
                            value: { amount: 2098.25, currencyCode: "AUD"},
                            percentage: 6.97,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "STW",
                        assetName: "Spdr S&P/Asx 200 Fund",
                        price: {amount: 76.15, currencyCode: "AUD"},
                        quantity: 347,
                        value: {amount: 26424.05, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 2905.20, currencyCode: "AUD"},
                            percentage: 25.65,
                        },
                        dividend:{
                            value: { amount: 1690.01, currencyCode: "AUD"},
                            percentage: 1.17,
                        },
                        currencyGain: {
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        totalReturn: {
                            value: { amount: 4595.21, currencyCode: "AUD"},
                            percentage: 26.82,
                        },
                        marketCode: "ASX",
                    },
                    {
                        assetSymbol: "NDQ",
                        assetName: "Betashares Nasdaq 100 ETF",
                        price: {amount: 51.18, currencyCode: "AUD"},
                        quantity: 356,
                        value: {amount: 18220.08, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 2452.72, currencyCode: "AUD"},
                            percentage: 15.56,
                        },
                        dividend:{
                            value: { amount: 9.82, currencyCode: "AUD"},
                            percentage: 0.06,
                        },
                        currencyGain: {
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        totalReturn: {
                            value: { amount: 2462.54, currencyCode: "AUD"},
                            percentage: 15.62,
                        },
                        marketCode: "ASX",
                    },
                    {
                        assetSymbol: "AVGO",
                        assetName: "Broadcom Inc",
                        price: {amount: 224.87, currencyCode: "USD"},
                        quantity: 36,
                        value: {amount: 12939.47, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 3381.19, currencyCode: "AUD"},
                            percentage: 37.42,
                        },
                        dividend: {
                            value: { amount: 69.63, currencyCode: "AUD"},
                            percentage: 0.77,
                        },
                        currencyGain: {
                            value: { amount: 513.15, currencyCode: "AUD"},
                            percentage: 5.68,
                        },
                        totalReturn: {
                            value: { amount: 3963.97, currencyCode: "AUD"},
                            percentage: 43.87,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "TSM",
                        assetName: "Taiwan Semiconductor Manufacturing Company Limited",
                        price: {amount: 206.12, currencyCode: "USD"},
                        quantity: 35,
                        value: {amount: 11531.10, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 2532.43, currencyCode: "AUD"},
                            percentage: 29.82,
                        },
                        dividend: {
                            value: { amount: 88.94, currencyCode: "AUD"},
                            percentage: 1.05,
                        },
                        currencyGain:{
                            value: { amount: 498.43, currencyCode: "AUD"},
                            percentage: 5.87,
                        },
                        totalReturn:{
                            value: { amount: 3119.80, currencyCode: "AUD"},
                            percentage: 36.74,
                        },
                        marketCode: "NYSE",
                    },
                    {
                        assetSymbol: "NVDA",
                        assetName: "Nvidia Corporation",
                        price: {amount: 129.84, currencyCode: "USD"},
                        quantity: 45,
                        value: {amount: 9317.17, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 904.51, currencyCode: "AUD"},
                            percentage: 11.24,
                        },
                        dividend: {
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0.02,
                        },
                        currencyGain: {
                            value: { amount: 382.77, currencyCode: "AUD"},
                            percentage: 4.76,
                        },
                        totalReturn: {
                            value: { amount: 1288.53, currencyCode: "AUD"},
                            percentage: 16.02,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "AMD",
                        assetName: "Adavanced Micro Devices Inc.",
                        price: {amount: 107.56, currencyCode: "USD"},
                        quantity: 39,
                        value: {amount: 6689.27, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: -2013.41, currencyCode: "AUD"},
                            percentage: -23.96,
                        },
                        dividend:{
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage:0,
                        },
                        currencyGain: {
                            value: { amount: 310.67, currencyCode: "AUD"},
                            percentage: 3.70,
                        },
                        totalReturn: {
                            value: { amount: -1702.74, currencyCode: "AUD"},
                            percentage: -20.26,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "MA",
                        assetName: "Mastercard Inc.",
                        price: {amount: 562.75, currencyCode: "USD"},
                        quantity: 7,
                        value: {amount: 6281.69, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 1599.11, currencyCode: "AUD"},
                            percentage: 35.88,
                        },
                        dividend:{
                            value: { amount: 29.74, currencyCode: "AUD"},
                            percentage: 0.67,
                        },
                        currencyGain: {
                            value: { amount: 236.93, currencyCode: "AUD"},
                            percentage: 5.32,
                        },
                        totalReturn: {
                            value: { amount: 1865.78, currencyCode: "AUD"},
                            percentage: 41.87,
                        },
                        marketCode: "NYSE",
                    },
                    {
                        assetSymbol: "UBER",
                        assetName: "Uber Technologies, Inc.",
                        price: {amount: 74.60, currencyCode: "USD"},
                        quantity: 52,
                        value: {amount: 6185.93, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 1097.22, currencyCode: "AUD"},
                            percentage: 21.39,
                        },
                        dividend:{
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        currencyGain: {
                            value: { amount: -29.90, currencyCode: "AUD"},
                            percentage: -0.58,
                        },
                        totalReturn: {
                            value: { amount: 1067.32, currencyCode: "AUD"},
                            percentage: 20.81,
                        },
                        marketCode: "NYSE",
                    },
                    {
                        assetSymbol: "FTNT",
                        assetName: "Fortinet Inc.",
                        price: {amount: 107.66, currencyCode: "USD"},
                        quantity: 34,
                        value: {amount: 5837.09, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 1751.43, currencyCode: "AUD"},
                            percentage: 45.87,
                        },
                        dividend:{
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        currencyGain: {
                            value: { amount: 277.12, currencyCode: "AUD"},
                            percentage: 7.26,
                        },
                        totalReturn: {
                            value: { amount: 2028.55, currencyCode: "AUD"},
                            percentage: 53.13,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "AMAT",
                        assetName: "Applied Materials Inc.",
                        price: {amount: 180.00, currencyCode: "USD"},
                        quantity: 15,
                        value: {amount: 4305.53, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: -95.22, currencyCode: "AUD"},
                            percentage: -2.33,
                        },
                        dividend:{
                            value: { amount: 9.40, currencyCode: "AUD"},
                            percentage: 0.23,
                        },
                        currencyGain: {
                            value: { amount: 315.70, currencyCode: "AUD"},
                            percentage: 7.71,
                        },
                        totalReturn: {
                            value: { amount: 229.88, currencyCode: "AUD"},
                            percentage: 5.62,
                        },
                        marketCode: "NASDAQ",
                    },
                    {
                        assetSymbol: "WTC",
                        assetName: "Wisetech Global Ltd",
                        price: {amount: 124.10, currencyCode: "AUD"},
                        quantity: 30,
                        value: {amount: 3723.00, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 1102.21, currencyCode: "AUD"},
                            percentage: 42.06,
                        },
                        dividend:{
                            value: { amount: 6.88, currencyCode: "AUD"},
                            percentage: 0.26,
                        },
                        currencyGain: {
                            value: { amount: 0, currencyCode: "AUD"},
                            percentage: 0,
                        },
                        totalReturn: {
                            value: { amount: 1109.09, currencyCode: "AUD"},
                            percentage: 42.32,
                        },
                        marketCode: "ASX",
                    },
                    {
                        assetSymbol: "GOOG",
                        assetName: "Alphabet Inc - Class C",
                        price: {amount: 187.14, currencyCode: "USD"},
                        quantity: 5,
                        value: {amount: 1492.11, currencyCode: "AUD"},
                        capitalGain: {
                            value: { amount: 502.15, currencyCode: "AUD"},
                            percentage: 28.05,
                        },
                        dividend:{
                            value: { amount: 4.57, currencyCode: "AUD"},
                            percentage: 0.28,
                        },
                        currencyGain: {
                            value: { amount: 54.54, currencyCode: "AUD"},
                            percentage: 3.31,
                        },
                        totalReturn: {
                            value: { amount: 561.26, currencyCode: "AUD"},
                            percentage: 31.05,
                        },
                        marketCode: "NASDAQ",
                    },

                ],
            },
        }};

    // Return mock data
    return HttpResponse.json(mockData)
};

export default getBalanceSummaryResolver;