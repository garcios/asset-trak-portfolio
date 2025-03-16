import AssetsTable from "./AssetsTable";
import React, {useEffect, useState} from "react";
import GraphQLService  from "../../services/graphql-service";
import {InvestmentLineItem} from "../../services/get-holdings-summary";
import {Card, CardContent, Typography} from "@mui/material";

const FAILED_TO_LOAD_MESSAGE = 'Failed to load balance summary';

function HoldingsContainer() {
    const [lineItems, setLineItems] = useState<InvestmentLineItem[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const loadBalanceSummary = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await GraphQLService.fetchHoldingsSummary('eb08df3c-958d-4ae8-b3ae-41ec04418786');
            setLineItems(data?.holdings);
        } catch (err) {
            setError(FAILED_TO_LOAD_MESSAGE);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadBalanceSummary();
    }, []);


    if (loading) return <p>Loading Holdings...</p>;
    if (error) return <p>Error: {error}</p>;

    return (
        <Card>
            <CardContent>
                <Typography variant="h6" component="div" fontWeight="bold" gutterBottom align="left">
                    Holdings
                </Typography>
                <AssetsTable items={lineItems} />
            </CardContent>
        </Card>
    );
}

export default HoldingsContainer;