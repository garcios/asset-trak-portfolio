import Grid from '@mui/material/Grid2';
import HoldingsContainer from "./HoldingsContainer";
import SummaryTotals from "./SummaryTotals";
import Performance from "./Performance";


const HoldingsSummary  = () => {
    return (
        <div>
            {/* Main Content Layout */}
            <Grid container spacing={2} sx={{ padding: 2 }}>

                {/* Summary */}
                <Grid  size={{ xs: 12 }}>
                   <SummaryTotals/>
                </Grid>

                {/* Portfolio Performance */}
                <Grid size={{ xs: 12}}>
                    <Performance/>
                </Grid>

                {/* Holdings */}
                <Grid size={{ xs: 12}}>
                   <HoldingsContainer/>
                </Grid>

            </Grid>
        </div>
    );
};

export default HoldingsSummary;
