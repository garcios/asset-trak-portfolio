import {graphql} from "msw";
import {getBalanceSummaryResolver} from "./getBalanceSummaryResolver";
import {BalanceSummaryResponse} from "../services/get-balance-summary";

const handlers = [
    graphql.query<BalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
];

export default handlers;