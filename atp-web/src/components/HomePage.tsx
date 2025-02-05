import React from "react";
import { AppBar, Toolbar, Typography, TextField, Card, CardContent, List, ListItem, ListItemText } from "@mui/material";
import Grid from '@mui/material/Grid2';// Import Grid2 from @mui/system"
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import BalanceSummary from "./BalanceSummary";
import BalanceHeader from "./BalanceHeader";

// TypeScript interfaces for stock data
interface StockData {
    time: string;
    price: number;
}

interface WatchlistStock {
    symbol: string;
    price: string;
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

// Sample watchlist data
const watchlistStocks: WatchlistStock[] = [
    { symbol: "GOOGL", price: "$204.02" },
    { symbol: "AMZN", price: "$237.68" },
    { symbol: "MSFT", price: "415.06" },
    { symbol: "AVGO", price: "$221.27" },
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

                {/* Stock Market Overview (Left Panel) */}
                <Grid size={{ xs: 12, md: 8 }}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6">Market Trends</Typography>
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

                {/* Watchlist (Right Panel) */}
                <Grid  size={{ xs: 12, md: 4 }}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6">Watchlist</Typography>
                            <List>
                                {watchlistStocks.map((stock, index) => (
                                    <ListItem key={index}>
                                        <ListItemText primary={stock.symbol} secondary={stock.price} />
                                    </ListItem>
                                ))}
                            </List>
                        </CardContent>
                    </Card>
                </Grid>

                {/* News Section */}
                <Grid size={{ xs: 12}}>
                    <Card>
                        <CardContent>
                           <BalanceHeader/>
                           <BalanceSummary/>
                        </CardContent>
                    </Card>
                </Grid>

            </Grid>
        </div>
    );
};

export default HomePage;
