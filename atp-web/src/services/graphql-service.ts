import client from './apollo-client';
import GET_BALANCE_SUMMARY, {BalanceSummaryResponse} from "./get-balance-summary";

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
}

export default new GraphQLService();
