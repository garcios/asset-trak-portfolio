import {Card, CardContent, Typography, useTheme} from "@mui/material";
import Grid from '@mui/material/Grid2';
import {FormatCurrency, FormatPercentage} from "../../utils/helper";
import React, {useEffect, useState} from "react";
import {SummaryTotalsType} from "../../services/get-summary-totals";
import GraphQLService from "../../services/graphql-service";


const FAILED_TO_LOAD_MESSAGE = 'Failed to load summary totals';


const SummaryTotals: React.FC = () => {
    const theme = useTheme();

    const [summaryTotals, setSummaryTotal] = useState<SummaryTotalsType>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const loadSummaryTotals = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await GraphQLService.fetchSummaryTotals('eb08df3c-958d-4ae8-b3ae-41ec04418786');
            setSummaryTotal(data.getSummaryTotals);
        } catch (err) {
            setError(FAILED_TO_LOAD_MESSAGE);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadSummaryTotals();
    }, []);


    if (loading) return <p>Loading Summary Totals...</p>;
    if (error) return <p>Error: {error}</p>;


    // Function to determine color based on positive/negative value
    const getColor = (value: number) =>
        value >= 0 ? theme.palette.success.main : theme.palette.error.main;

    return (
        <Card>
            <CardContent>
                <Typography variant="h6" fontWeight="bold" gutterBottom align="left">
                    Summary
                </Typography>


                <Grid container spacing={2}>
                    {/* Portfolio Value*/}
                    <Grid size={{ xs: 2 }} >
                        <Typography variant="body1" fontWeight="bold">
                            Portfolio Value
                        </Typography>
                        <Typography variant="h4" fontWeight="bold">{FormatCurrency(summaryTotals.portfolioValue.amount,
                            summaryTotals.portfolioValue.currencyCode)} </Typography>
                    </Grid>

                    {/* Capital Gain */}
                    <Grid  size={{ xs: 2}}>
                        <Typography variant="body1" fontWeight="bold">
                            Capital Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(summaryTotals.capitalGain.amount, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(summaryTotals.capitalGain.percentage)}>
                            {FormatPercentage(summaryTotals.capitalGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Dividends */}
                    <Grid size={{ xs: 2}}>
                        <Typography variant="body1" fontWeight="bold">
                            Dividends
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(summaryTotals.dividends.amount, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(summaryTotals.dividends.percentage)}>
                            {FormatPercentage(summaryTotals.dividends.percentage)}
                        </Typography>
                    </Grid>

                    {/* Currency Gain */}
                    <Grid size={{ xs: 2 }} >
                        <Typography variant="body1" fontWeight="bold">
                            Currency Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(summaryTotals.currencyGain.amount,"AUD")}</Typography>
                        <Typography variant="body2" color={getColor(summaryTotals.currencyGain.percentage)}>
                            {FormatPercentage(summaryTotals.currencyGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Total Return */}
                    <Grid size={{ xs: 2 }} >
                        <Typography variant="body1" fontWeight="bold">
                            Total Return
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(summaryTotals.totalReturn.amount,"AUD")}</Typography>
                        <Typography variant="body2" color={getColor(summaryTotals.totalReturn.percentage)}>
                            {FormatPercentage(summaryTotals.totalReturn.percentage)}
                        </Typography>
                    </Grid>

                </Grid>
            </CardContent>
        </Card>
    );
};


export default SummaryTotals;