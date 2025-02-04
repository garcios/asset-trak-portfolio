import * as React from 'react';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid2';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import BalanceSummary from "../components/BalanceSummary"; // Import Typography

const Item = styled(Paper)(({ theme }) => ({
    padding: theme.spacing(1),
    textAlign: 'left', // Align text to the left
    color: theme.palette.text.secondary,
    height: '100%', // Make sure paper takes full height of grid cell for consistent look
    display: 'flex',
    alignItems: 'center', // Vertically center content
    paddingLeft: theme.spacing(1), // Add some left padding for better readability
}));

export default function MasterContainer(props: any)  {
    return (
        <Box sx={{ flexGrow: 1, padding: 1 }}>
            <Grid container spacing={1}>
                <Grid size={{ xs: 12 }}>
                    <Item elevation={2}> {/* Added elevation for a more prominent header */}
                        <Typography variant="h5" component="h1"> {/* Use Typography for header */}
                            Asset Tracking
                        </Typography>
                    </Item>
                </Grid>

                <Grid size={{ xs: 12, md: 2 }}>
                    <Grid container spacing={1} direction="column"> {/* Nested grid for the sidebar */}
                        <Grid size={{ xs: 12 }}>
                            <Item elevation={1}>Overview</Item>
                        </Grid>
                        <Grid  size={{ xs: 12 }}>
                            <Item elevation={1}>Watchlist</Item>
                        </Grid>
                        <Grid  size={{ xs: 12 }}>
                            <Item elevation={1}>News</Item>
                        </Grid>
                        <Grid  sx={{ flex: 1 }} >
                            <VerticalStretchGrid/>
                        </Grid>
                    </Grid>
                </Grid>

                <Grid size={{ xs: 12, md: 10 }}>
                    <Grid container spacing={1} direction="column"> {/* Nested grid for the main content */}
                        <Grid size={{ xs: 12 }}>
                            <Item elevation={1}>Portfolio Value: $321,000</Item>
                        </Grid>
                        <Grid sx={{ flex: 1 }}>
                            <BalanceSummaryTable/>
                        </Grid>
                    </Grid>
                </Grid>

                {/* Footer (Optional) */}
                <Grid size={{ xs: 12 }}>
                    <Item elevation={1}>
                    </Item>
                </Grid>
            </Grid>
        </Box>
    );
}

function VerticalStretchGrid() {
    return (
        <Box sx={{ height: "50vh", display: "flex", flexDirection: "column" }}>
            <Grid container sx={{ flex: 1 }}>
                <Grid size={{ xs: 12 }} sx={{ backgroundColor: 'background.default', height: "100%" }}>
                    Left Content
                </Grid>
            </Grid>
        </Box>
    );
}

function BalanceSummaryTable() {
    return (
        <Box sx={{ height: "50vh", display: "flex", flexDirection: "column" }}>
            <Grid container sx={{ flex: 1 }}>
                <Grid size={{ xs: 12 }} sx={{ backgroundColor: 'background.default', height: "100%" }}>
                    <Item elevation={1}><BalanceSummary/></Item>
                </Grid>
            </Grid>
        </Box>
    );
}