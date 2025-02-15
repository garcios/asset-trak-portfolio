import {GraphQLResponseResolver, HttpResponse} from "msw";

const getPerformanceNumbersResolver:GraphQLResponseResolver = ({query, variables}) => {
    const mockData = { data: {
            getHistoricalValues: [
                { tradeDate: "Jan 2020", amount: 10000, currencyCode: "AUD" },
                { tradeDate: "Jul 2020", amount: 20000, currencyCode: "AUD" },
                { tradeDate: "Jan 2021", amount: 80000, currencyCode: "AUD" },
                { tradeDate: "Jul 2021", amount: 70000, currencyCode: "AUD" },
                { tradeDate: "Jan 2022", amount: 100000, currencyCode: "AUD"},
                { tradeDate: "Jul 2022", amount: 110000, currencyCode: "AUD" },
                { tradeDate: "Jan 2023", amount: 140000, currencyCode: "AUD" },
                { tradeDate: "Jul 2023", amount: 200000, currencyCode: "AUD"},
                { tradeDate: "Jan 2024", amount: 280000, currencyCode: "AUD" },
                { tradeDate: "Jul 2024", amount: 300000, currencyCode: "AUD" },
                { tradeDate: "Jan 2025", amount: 316000, currencyCode: "AUD" },
            ]
        }};

    return HttpResponse.json(mockData)
}

export default getPerformanceNumbersResolver;