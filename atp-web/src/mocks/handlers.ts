import {graphql} from "msw";
import {BalanceSummaryResponse} from "../services/get-balance-summary";
import getBalanceSummaryResolver from "./get-balance-summary-resolver";
import getSummaryTotalsResolver from "./get-summary-totals-resolver";
import {SummaryTotalsResponse} from "../services/get-summary-totals";

const handlers = [
    graphql.query<BalanceSummaryResponse, { accountId: string }>(
        'GetBalanceSummary',
        getBalanceSummaryResolver,
    ),
    graphql.query<SummaryTotalsResponse, { accountId: string }>(
        'GetSummaryTotals',
        getSummaryTotalsResolver,
    ),
];

export default handlers;