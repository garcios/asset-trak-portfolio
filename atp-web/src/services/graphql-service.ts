import client from './apollo-client';
import GET_HOLDINGS_SUMMARY, {HoldingsSummaryResponse, InvestmentLineItem} from "./get-holdings-summary";
import GET_SUMMARY_TOTALS, {SummaryTotalsResponse} from "./get-summary-totals";
import GET_PERFORMANCE_NUMBERS, {PerformanceDataResponse} from "./get-performance-numbers";

class GraphQLService {
    async fetchHoldingsSummary(accountId: string): Promise<HoldingsSummaryResponse> {
        const variables = { accountId: accountId };

        try {
            const { data: holdingsSummary } = await client.query<HoldingsSummaryResponse>({
                query: GET_HOLDINGS_SUMMARY,
                variables
            });

            return holdingsSummary;
        } catch (error) {
            console.error(`Error fetching balance summaries for accountId: ${accountId}`, error);
            throw error;
        }
    }

    async fetchSummaryTotals(accountId: string): Promise<SummaryTotalsResponse> {
        const variables = { accountId: accountId };

        try {
            const { data: summaryTotals } = await client.query<SummaryTotalsResponse>({
                query: GET_SUMMARY_TOTALS,
                variables
            });

            return summaryTotals;
        } catch (error) {
            console.error(`Error fetching summary totals for accountId: ${accountId}`, error);
            throw error;
        }
    }

    async fetchPerformanceNumbers(accountId: string): Promise<PerformanceDataResponse> {
        const variables = { accountId: accountId, startDate: '2020-07-29', endDate: '2025-03-10' };

        try {
            const { data: performanceNumbers } = await client.query<PerformanceDataResponse>({
                query: GET_PERFORMANCE_NUMBERS,
                variables
            });

            return performanceNumbers;
        } catch (error) {
            console.error(`Error fetching summary totals for accountId: ${accountId}`, error);
            throw error;
        }
    }
}

export default new GraphQLService();
