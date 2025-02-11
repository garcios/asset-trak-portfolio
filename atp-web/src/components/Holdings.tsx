import { Typography,  Card, CardContent} from "@mui/material";
import Grid from '@mui/material/Grid2';
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import BalanceSummary from "./BalanceSummary";
import SummaryTotals from "./SummaryTotals";

// TypeScript interfaces for stock data
interface PerformanceData {
    date: string;
    value: number;
}

// Sample stock price data for chart
const samplePerformanceData: PerformanceData[] = [
    { date: "Jan 2020", value: 10000 },
    { date: "Jul 2020", value: 20000 },
    { date: "Jan 2021", value: 80000 },
    { date: "Jul 2021", value: 70000 },
    { date: "Jan 2022", value: 100000 },
    { date: "Jul 2022", value: 110000 },
    { date: "Jan 2023", value: 140000 },
    { date: "Jul 2023", value: 200000 },
    { date: "Jan 2024", value: 280000 },
    { date: "Jul 2024", value: 300000 },
    { date: "Jan 2025", value: 316000 },
];

const Holdings  = () => {
    return (
        <div>
            {/* Main Content Layout */}
            <Grid container spacing={2} sx={{ padding: 2 }}>

                {/* Summary */}
                <Grid  size={{ xs: 12 }}>
                    <Card>
                        <CardContent>
                            <SummaryTotals/>
                        </CardContent>
                    </Card>
                </Grid>

                {/* Portfolio Performance */}
                <Grid size={{ xs: 12}}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6" fontWeight="bold" gutterBottom align="left">Performance</Typography>
                            <ResponsiveContainer width="100%" height={300}>
                                <LineChart data={samplePerformanceData}>
                                    <XAxis dataKey="date" />
                                    <YAxis dataKey="value" />
                                    <Tooltip />
                                    <Line type="monotone" dataKey="value" stroke="#1976d2" />
                                </LineChart>
                            </ResponsiveContainer>
                        </CardContent>
                    </Card>
                </Grid>


                {/* Holdings */}
                <Grid size={{ xs: 12}}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6" component="div" fontWeight="bold" gutterBottom align="left">
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

export default Holdings;
