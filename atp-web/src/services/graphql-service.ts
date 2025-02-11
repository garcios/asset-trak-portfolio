import client from './apollo-client';
import GET_BALANCE_SUMMARY, {BalanceSummaryResponse} from "./get-balance-summary";

class GraphQLService {
    async getBalanceSummaries(accountId:string): Promise<BalanceSummaryResponse> {
        try {
            const {data} = await client.query<BalanceSummaryResponse>({
                query: GET_BALANCE_SUMMARY,
                variables: {
                    accountId: accountId
                }
            });
            console.log('getBalanceSummaries:',data);

            return data;
        } catch (error) {
            console.error('Error fetching balance summaries:', error);
            throw error;
        }
    }
}

export default new GraphQLService();
