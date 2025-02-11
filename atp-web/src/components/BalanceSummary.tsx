import AssetsTable from "./AssetsTable";
import React, {useEffect, useState} from "react";
import GraphQLService  from "../services/graphql-service";
import {LineItem} from "../services/get-balance-summary";

function BalanceSummary() {
    const [lineItems, setLineItems] = useState<LineItem[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchCountries = async () => {
            try {
                const data = await GraphQLService.getBalanceSummaries('eb08df3c-958d-4ae8-b3ae-41ec04418786');
                setLineItems(data.balanceItems);
            } catch (err) {
                setError('Failed to load countries');
            } finally {
                setLoading(false);
            }
        };

        fetchCountries().catch(
            (err) => setError('Failed to load countries'))
        return () => setLoading(true);

    }, []);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error}</p>;

    return (
        <>
            <AssetsTable items={lineItems} />
        </>

    );
}

export default BalanceSummary;