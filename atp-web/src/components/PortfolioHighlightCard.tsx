import React from "react";
import {Box, Card, CardContent, Typography, useTheme} from "@mui/material";
import Grid from '@mui/material/Grid2';
import {FormatCurrency, FormatPercentage} from "../utils/helper";
import Paper from "@mui/material/Paper";
import ArrowDropUpIcon from "@mui/icons-material/ArrowDropUp";
import ArrowDropDownIcon from "@mui/icons-material/ArrowDropDown";

// TypeScript interface for the component props
interface PortfolioHighlightProps {
    capitalGain: { value: number; percentage: number };
    dividends: { value: number; percentage: number };
    currencyGain: { value: number; percentage: number };
    totalReturn: { value: number; percentage: number };
}

const PortfolioHighlightCard: React.FC<PortfolioHighlightProps> = ({
                                                                       capitalGain,
                                                                       dividends,
                                                                       currencyGain,
                                                                       totalReturn,
                                                                   }) => {
    const theme = useTheme();

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
                        <Typography variant="h4" fontWeight="bold">{FormatCurrency(316619.50,"AUD")} </Typography>
                    </Grid>

                    {/* Capital Gain */}
                    <Grid  size={{ xs: 2}}>
                        <Typography variant="body1" fontWeight="bold">
                            Capital Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(capitalGain.value, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(capitalGain.percentage)}>
                            {FormatPercentage(capitalGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Dividends */}
                    <Grid size={{ xs: 2}}>
                        <Typography variant="body1" fontWeight="bold">
                            Dividends
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(dividends.value, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(dividends.percentage)}>
                            {FormatPercentage(dividends.percentage)}
                        </Typography>
                    </Grid>

                    {/* Currency Gain */}
                    <Grid size={{ xs: 2 }} >
                        <Typography variant="body1" fontWeight="bold">
                            Currency Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(currencyGain.value,"AUD")}</Typography>
                        <Typography variant="body2" color={getColor(currencyGain.percentage)}>
                            {FormatPercentage(currencyGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Total Return */}
                    <Grid size={{ xs: 2 }} >
                        <Typography variant="body1" fontWeight="bold">
                            Total Return
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(totalReturn.value,"AUD")}</Typography>
                        <Typography variant="body2" color={getColor(totalReturn.percentage)}>
                            {FormatPercentage(currencyGain.percentage)}
                        </Typography>
                    </Grid>

                </Grid>
            </CardContent>
        </Card>
    );
};


export default PortfolioHighlightCard;