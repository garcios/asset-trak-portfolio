import client from './apollo-client';
import GET_BALANCE_SUMMARY, {BalanceSummaryResponse} from "./get-balance-summary";
import GET_SUMMARY_TOTALS, {SummaryTotalsResponse} from "./get-summary-totals";

class GraphQLService {
    async fetchBalanceSummaries(accountId: string): Promise<BalanceSummaryResponse> {
        const variables = { id: accountId };

        try {
            const { data: balanceSummary } = await client.query<BalanceSummaryResponse>({
                query: GET_BALANCE_SUMMARY,
                variables
            });

            return balanceSummary;
        } catch (error) {
            console.error(`Error fetching balance summaries for accountId: ${accountId}`, error);
            throw error;
        }
    }

    async fetchSummaryTotals(accountId: string): Promise<SummaryTotalsResponse> {
        const variables = { id: accountId };

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

}

export default new GraphQLService();
