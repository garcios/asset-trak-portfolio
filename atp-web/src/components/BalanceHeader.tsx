import Typography from "@mui/material/Typography";
import {gql, useQuery} from "@apollo/client";
import {FormatCurrency} from "../utils/helper";

const GET_BALANCE_TOTAL = gql`
    query {
        getBalanceSummary(accountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786") {
            accountId
            totalValue {
                amount
                currencyCode
            }
        }
    }
`;

function BalanceHeader() {
    const {loading, error, data} = useQuery(GET_BALANCE_TOTAL);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    const {amount, currencyCode} = data?.getBalanceSummary?.totalValue;

    return (
        <>
            <Typography variant="h5" component="div" align="left">
               Total Value: {FormatCurrency(amount, currencyCode)}
            </Typography>
        </>
    )
}

export default BalanceHeader;