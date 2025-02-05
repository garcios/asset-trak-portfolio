import React from "react";
import { AppBar, Toolbar, Typography, TextField, Card, CardContent, List, ListItem, ListItemText } from "@mui/material";
import Grid from '@mui/material/Grid2';
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import BalanceSummary from "./BalanceSummary";
import PortfolioHighlightCard from "./PortfolioHighlightCard";

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
            {/* Top Navigation Bar */}
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" sx={{ flexGrow: 1 }}>
                        Asset Performance Tracking
                    </Typography>
                    <TextField
                        variant="outlined"
                        size="small"
                        placeholder="Search stocks..."
                        sx={{ backgroundColor: "white", borderRadius: 1 }}
                    />
                </Toolbar>
            </AppBar>

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
                                capitalGain={{ value: 5000, percentage: 8.5 }}
                                dividends={{ value: 1200, percentage: 2.5 }}
                                currencyGain={{ value: -300, percentage: -0.8 }}
                                totalReturn={{ value: 68000, percentage: 10.2 }}
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
