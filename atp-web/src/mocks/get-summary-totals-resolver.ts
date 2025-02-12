import {GraphQLResponseResolver, HttpResponse} from "msw";

const getSummaryTotalsResolver:GraphQLResponseResolver = ({query, variables}) => {
    const mockData = { data: {
            summaryTotals: {
                portfolioValue: {
                    amount: 316076.18,
                    currencyCode: "AUD"
                },
                capitalGain: {
                    amount: 56287.88,
                    currencyCode: "AUD",
                    percentage: 19.56
                },
                dividends:{
                    amount: 8192.71,
                    currencyCode: "AUD",
                    percentage: 2.85
                },
                currencyGain: {
                    amount: 7205.00,
                    currencyCode: "AUD",
                    percentage: 2.5
                },
                totalReturn: {
                    amount: 71685.59,
                    currencyCode: "AUD",
                    percentage: 24.91
                }
            }
        }};

    return HttpResponse.json(mockData)
}

export default getSummaryTotalsResolver;