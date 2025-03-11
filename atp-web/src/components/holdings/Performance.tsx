import React, {useEffect, useState} from "react";
import {Typography,Card, CardContent} from "@mui/material";
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from "recharts";
import GraphQLService from "../../services/graphql-service";
import {PerformanceData} from "../../services/get-performance-numbers";
import { format } from 'date-fns';
import { format as format2 }  from 'd3-format';


const FAILED_TO_LOAD_MESSAGE = 'Failed to load performance values';

const CustomizedXAxisTick = ({ x, y, payload }) => {
    // Customize the date format here using date-fns
    const formattedDate = format(payload.value, 'MMM yyyy'); // Example format

    return (
        <g transform={`translate(${x},${y})`}>
            <text x={0} y={0} dy={10} textAnchor="end" fill="#666" transform="rotate(-35)">
                {formattedDate}
            </text>
        </g>
    );
};

const CustomizedYAxisTick = ({ x, y, payload }) => {
    const formattedValue = format2('$,')(payload.value); // Format with comma and dollar sign

    return (
        <g transform={`translate(${x},${y})`}>
            <text x={0} y={0} dy={0} textAnchor="end" fill="#666">
                {formattedValue}
            </text>
        </g>
    );
};


const Performance: React.FC = () => {
    const [performanceNumbers, setPerformanceNumbers] = useState<PerformanceData[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const loadPerformanceNumbers = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await GraphQLService.fetchPerformanceNumbers('eb08df3c-958d-4ae8-b3ae-41ec04418786');
            setPerformanceNumbers(data.getPerformanceHistory);
        } catch (err) {
            setError(FAILED_TO_LOAD_MESSAGE);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadPerformanceNumbers();
    }, []);


    if (loading) return <p>Loading Performance Chart...</p>;
    if (error) return <p>Error: {error}</p>;


    return (
        <Card>
            <CardContent>
                <Typography variant="h6" fontWeight="bold" gutterBottom align="left">Performance</Typography>
                <ResponsiveContainer width="100%" height={400}>
                    <LineChart data={performanceNumbers} margin={{ top: 5, right: 30, left: 20, bottom: 30 }}>
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="tradeDate"  tick={<CustomizedXAxisTick/>} />
                        <YAxis tick={<CustomizedYAxisTick/>}/>
                        <Tooltip />
                        <Legend verticalAlign="top"/>
                        <Line type="monotone" dataKey="value" stroke="#8884d8" strokeWidth={2} />
                        <Line type="monotone" dataKey="cost" stroke="#82ca9d" strokeDasharray="5 5"/>
                    </LineChart>
                </ResponsiveContainer>
            </CardContent>
        </Card>
    );
};

export default Performance;
