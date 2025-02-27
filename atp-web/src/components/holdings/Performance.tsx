import React, {useEffect, useState} from "react";
import {Typography,Card, CardContent} from "@mui/material";
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from "recharts";
import GraphQLService from "../../services/graphql-service";
import {PerformanceData} from "../../services/get-performance-numbers";

const FAILED_TO_LOAD_MESSAGE = 'Failed to load performance values';

const Performance: React.FC = () => {
    const [performanceNumbers, setPerformanceNumbers] = useState<PerformanceData[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const loadPerformanceNumbers = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await GraphQLService.fetchPerformanceNumbers('eb08df3c-958d-4ae8-b3ae-41ec04418786');
            setPerformanceNumbers(data.getHistoricalValues);
        } catch (err) {
            setError(FAILED_TO_LOAD_MESSAGE);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadPerformanceNumbers();
    }, []);


    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error}</p>;


    return (
        <Card>
            <CardContent>
                <Typography variant="h6" fontWeight="bold" gutterBottom align="left">Performance</Typography>
                <ResponsiveContainer width="100%" height={300}>
                    <LineChart data={performanceNumbers}>
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="tradeDate"/>
                        <YAxis/>
                        <Tooltip />
                        <Legend />
                        <Line type="monotone" dataKey="value" stroke="#8884d8" />
                        <Line type="monotone" dataKey="cost" stroke="#82ca9d" strokeDasharray="5 5"/>
                    </LineChart>
                </ResponsiveContainer>
            </CardContent>
        </Card>
    );
};

export default Performance;
