import {GraphQLResponseResolver, HttpResponse} from "msw";

const getPerformanceNumbersResolver:GraphQLResponseResolver = ({query, variables}) => {
    const mockData = { data: {
            dataItems: [
                { date: "Jan 2020", value: 10000 },
                { date: "Jul 2020", value: 20000 },
                { date: "Jan 2021", value: 80000 },
                { date: "Jul 2021", value: 70000 },
                { date: "Jan 2022", value: 100000 },
                { date: "Jul 2022", value: 110000 },
                { date: "Jan 2023", value: 140000 },
                { date: "Jul 2023", value: 200000 },
                { date: "Jan 2024", value: 280000 },
                { date: "Jul 2024", value: 300000 },
                { date: "Jan 2025", value: 316000 },
            ]
        }};

    return HttpResponse.json(mockData)
}

export default getPerformanceNumbersResolver;