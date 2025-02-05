import React from "react";
import { AppBar, Toolbar, Typography, TextField, Card, CardContent, List, ListItem, ListItemText } from "@mui/material";
import Grid from '@mui/material/Grid2';
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import BalanceSummary from "./BalanceSummary";
import PortfolioHighlightCard from "./PortfolioHighlightCard";
import TopNavBar from "./TopNavBar";

// TypeScript interfaces for stock data
interface StockData {
    time: string;
    price: number;
}

// Sample stock price data for chart
const sampleStockData: StockData[] = [
    { time: "9 AM", price: 100 },
    { time: "10 AM", price: 155 },
    { time: "11 AM", price: 160 },
    { time: "12 PM", price: 150 },
    { time: "1 PM", price: 110 },
    { time: "2 PM", price: 160 },
];

const HomePage: React.FC = () => {
    return (
        <div>
            {/* Main Content Layout */}
            <Grid container spacing={2} sx={{ padding: 2 }}>

                {/* Portfolio Performance (Left Panel) */}
                <Grid size={{ xs: 12, md: 8 }}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6">Performance</Typography>
                            <ResponsiveContainer width="100%" height={300}>
                                <LineChart data={sampleStockData}>
                                    <XAxis dataKey="time" />
                                    <YAxis />
                                    <Tooltip />
                                    <Line type="monotone" dataKey="price" stroke="#1976d2" />
                                </LineChart>
                            </ResponsiveContainer>
                        </CardContent>
                    </Card>
                </Grid>

                {/* Highlights (Right Panel) */}
                <Grid  size={{ xs: 12, md: 4 }}>
                    <Card>
                        <CardContent>
                            <PortfolioHighlightCard
                                capitalGain={{ value: 53421.81, percentage: 20.88 }}
                                dividends={{ value: 3038.29, percentage: 1.19 }}
                                currencyGain={{ value: 7079.06, percentage: 2.77 }}
                                totalReturn={{ value: 63539.16, percentage: 24.84 }}
                            />
                        </CardContent>
                    </Card>
                </Grid>

                {/* Holdings */}
                <Grid size={{ xs: 12}}>
                    <Card>
                        <CardContent>
                            <Typography variant="h5" component="div" align="left">
                                Holdings
                            </Typography>
                           <BalanceSummary/>
                        </CardContent>
                    </Card>
                </Grid>

            </Grid>
        </div>
    );
};

export default HomePage;
