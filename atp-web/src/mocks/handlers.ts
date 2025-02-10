import {graphql, GraphQLResponseResolver, HttpResponse} from "msw";
import {getBalanceSummaryResolver, GetBalanceSummaryResponse} from "./getBalanceSummaryResolver";


const handlers = [
    graphql.query<GetBalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
];

export default handlers;