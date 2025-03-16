import {graphql} from "msw";
import {HoldingsSummaryResponse} from "../services/get-holdings-summary";
import getHoldingsSummaryResolver from "./get-holdings-summary-resolver";
import getSummaryTotalsResolver from "./get-summary-totals-resolver";
import {SummaryTotalsResponse} from "../services/get-summary-totals";
import {PerformanceDataResponse} from "../services/get-performance-numbers";
import getPerformanceNumbersResolver from "./get-performance-numbers-resolvers";

const handlers = [
    graphql.query<HoldingsSummaryResponse, { accountId: string }>(
        'holdings',
        getHoldingsSummaryResolver,
    ),
    graphql.query<SummaryTotalsResponse, { accountId: string }>(
        'summaryTotals',
        getSummaryTotalsResolver,
    ),
    graphql.query<PerformanceDataResponse, { accountId: string }>(
        'performanceNumbers',
        getPerformanceNumbersResolver,
    ),
];

export default handlers;