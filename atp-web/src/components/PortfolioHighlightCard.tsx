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
                <Typography variant="h6" fontWeight="bold" gutterBottom>
                    Portfolio Highlights
                </Typography>

                <Grid container spacing={2}>
                    {/* Capital Gain */}
                    <Grid  size={{ xs: 6}}>
                        <Typography variant="body1" fontWeight="bold">
                            Capital Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(capitalGain.value, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(capitalGain.percentage)}>
                            {FormatPercentage(capitalGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Dividends */}
                    <Grid size={{ xs: 6}}>
                        <Typography variant="body1" fontWeight="bold">
                            Dividends
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(dividends.value, "AUD")}</Typography>
                        <Typography variant="body2" color={getColor(dividends.percentage)}>
                            {FormatPercentage(dividends.percentage)}
                        </Typography>
                    </Grid>

                    {/* Currency Gain */}
                    <Grid size={{ xs: 6 }} sx={{ mt: 2 }}>
                        <Typography variant="body1" fontWeight="bold">
                            Currency Gain
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(currencyGain.value,"AUD")}</Typography>
                        <Typography variant="body2" color={getColor(currencyGain.percentage)}>
                            {FormatPercentage(currencyGain.percentage)}
                        </Typography>
                    </Grid>

                    {/* Portfolio Value*/}
                    <Grid size={{ xs: 6 }} sx={{ mt: 2 }}>
                        <Typography variant="body1" fontWeight="bold">
                            Portfolio Value
                        </Typography>
                        <Typography variant="h6">{FormatCurrency(316619.50,"AUD")}</Typography>
                    </Grid>

                    {/* Total Return */}
                    <Grid size={{ xs: 12 }} sx={{ mt: 2 }}>
                        <Paper
                            sx={{
                                backgroundColor: getColor(totalReturn.percentage),
                                color: "white",
                                padding: 2,
                                borderRadius: 2,
                                textAlign: "center",
                                boxShadow: 3,
                            }}
                        >
                            <Typography variant="h6" fontWeight="bold">
                                Total Return
                            </Typography>
                            <Typography variant="h4" fontWeight="bold">
                                {FormatCurrency(totalReturn.value,"AUD")}
                            </Typography>
                            <Box display="flex" justifyContent="center" alignItems="center">
                                {totalReturn.percentage >= 0 ? <ArrowDropUpIcon fontSize="large" /> : <ArrowDropDownIcon fontSize="large" />}
                                <Typography variant="h6" fontWeight="bold">
                                    {FormatPercentage(totalReturn.percentage)}
                                </Typography>
                            </Box>
                        </Paper>
                    </Grid>
                </Grid>
            </CardContent>
        </Card>
    );
};


export default PortfolioHighlightCard;