import {GraphQLResponseResolver, HttpResponse} from "msw";

const getPerformanceNumbersResolver:GraphQLResponseResolver = ({query, variables}) => {
    const mockData = { data: {
            performanceHistory: [
                { tradeDate: "Jan 2020", cost: 10000, value: 10000, currencyCode: "AUD" },
                { tradeDate: "Jul 2020", cost: 15000, value: 20000, currencyCode: "AUD" },
                { tradeDate: "Jan 2021", cost: 16000, value: 80000, currencyCode: "AUD" },
                { tradeDate: "Jul 2021", cost: 60000, value: 70000, currencyCode: "AUD" },
                { tradeDate: "Jan 2022", cost: 80000, value: 100000, currencyCode: "AUD"},
                { tradeDate: "Jul 2022", cost: 121000, value: 110000, currencyCode: "AUD" },
                { tradeDate: "Jan 2023", cost: 16000, value: 140000, currencyCode: "AUD" },
                { tradeDate: "Jul 2023", cost: 50000, value: 200000, currencyCode: "AUD"},
                { tradeDate: "Jan 2024", cost: 100000, value: 280000, currencyCode: "AUD" },
                { tradeDate: "Jul 2024", cost: 150000, value: 300000, currencyCode: "AUD" },
                { tradeDate: "Jan 2025", cost: 200000, value: 316000, currencyCode: "AUD" },
            ]
        }};

    return HttpResponse.json(mockData)
}

export default getPerformanceNumbersResolver;