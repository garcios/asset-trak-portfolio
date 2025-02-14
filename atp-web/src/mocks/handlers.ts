import {graphql} from "msw";
import {BalanceSummaryResponse} from "../services/get-balance-summary";
import getBalanceSummaryResolver from "./get-balance-summary-resolver";
import getSummaryTotalsResolver from "./get-summary-totals-resolver";
import {SummaryTotalsResponse} from "../services/get-summary-totals";
import {PerformanceDataResponse} from "../services/get-performance-numbers";
import getPerformanceNumbersResolver from "./get-performance-numbers-resolvers";

const handlers = [
    graphql.query<BalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
    graphql.query<SummaryTotalsResponse, { accountId: string }>(
        'GetSummaryTotals',
        getSummaryTotalsResolver,
    ),
    graphql.query<PerformanceDataResponse, { accountId: string }>(
        'GetPerformanceNumbers',
        getPerformanceNumbersResolver,
    ),
];

export default handlers;