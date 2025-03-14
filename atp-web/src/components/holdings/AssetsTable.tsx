import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import {Typography, useTheme} from "@mui/material";
import {FormatCurrency, FormatPercentage} from "../../utils/helper";
import React from "react";
import {InvestmentLineItem} from "../../services/get-holdings-summary";

interface Column {
    id: 'assetName' |  'price' | 'currency' | 'quantity' | 'value' | 'capitalGain' | 'dividend' | 'currencyGain' | 'totalReturn';
    label: string;
    minWidth?: number;
    align?: 'right' | 'left' | 'center' | 'justify' | 'inherit';
}

const columns: readonly Column[] = [
    {
        id: 'assetName',
        label: 'Name',
        minWidth: 170 },
    {
        id: 'price',
        label: 'Price',
        minWidth: 100,
        align: 'left',
    },
    {
        id: 'quantity',
        label: 'Weight & Quantity',
        minWidth: 100,
    },
    {
        id: 'value',
        label: 'Value & Cost',
        minWidth: 100,
        align: 'left',
    },
    {
        id: 'capitalGain',
        label: 'Capital Gain',
        minWidth: 100,
        align: 'left',
    },
    {
        id: 'dividend',
        label: 'Dividend',
        minWidth: 100,
        align: 'left',
    },
    {
        id: 'currencyGain',
        label: 'Currency Gain',
        minWidth: 100,
        align: 'left',
    },
    {
        id: 'totalReturn',
        label: 'Total Return',
        minWidth: 100,
        align: 'left',
    },
];

export default function AssetsTable({ items }: { items: InvestmentLineItem[] }) {
    const theme = useTheme();

    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(10);

    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    // Function to determine color based on positive/negative value
    const getColor = (value: number) =>
        value >= 0 ? theme.palette.success.main : theme.palette.error.main;

    return (
        <Paper sx={{ width: '100%', overflow: 'hidden' }}>
            <TableContainer sx={{ maxHeight: 440 }}>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            {columns.map((column) => (
                                <TableCell
                                    key={column.id}
                                    align={column.align}
                                    style={{ minWidth: column.minWidth }}
                                >
                                    {column.label}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {items
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((row) => {
                                return (
                                    <TableRow hover role="checkbox" tabIndex={-1} key={row.assetSymbol}  >
                                         <TableCell key='ticker' align="left">
                                            <Typography variant="body1">{row.assetSymbol}.{row.marketCode}</Typography>
                                            <Typography variant="body2" color="textSecondary">{row.assetName}</Typography>
                                         </TableCell>
                                        <TableCell key='price' align="left">
                                            <Typography variant="body1">{FormatCurrency(row.price?.amount, row.price?.currencyCode)}</Typography>
                                        </TableCell>
                                        <TableCell key='quantity' align="left">
                                            <Typography variant="body1">{FormatPercentage(row.weight)}</Typography>
                                            <Typography variant="body2" color="textSecondary">{row.quantity}</Typography>
                                        </TableCell>
                                        <TableCell key='value' align="left">
                                            <Typography variant="body1">{FormatCurrency(row.value?.amount, row.value?.currencyCode)}</Typography>
                                            <Typography variant="body2" color="textSecondary">{FormatCurrency(row.cost?.amount, row.cost?.currencyCode)}</Typography>
                                        </TableCell>
                                        <TableCell key='capitalGain' align="left">
                                            <Typography
                                                variant="body1"
                                                color={getColor(row.capitalGain.amount)}>{FormatCurrency(row.capitalGain?.amount, row.capitalGain?.currencyCode)}
                                            </Typography>
                                            <Typography
                                                sx={{ opacity: 0.7 }}
                                                variant="body2"
                                                color={getColor(row.capitalGain.percentage)}>{FormatPercentage(row.capitalGain.percentage)}
                                            </Typography>
                                        </TableCell>
                                        <TableCell key='dividend' align="left">
                                            <Typography
                                                variant="body1"
                                                color={getColor(row.dividend.amount)}>{FormatCurrency(row.dividend?.amount, row.dividend?.currencyCode)}
                                            </Typography>
                                            <Typography
                                                sx={{ opacity: 0.7 }}
                                                variant="body2"
                                                color={getColor(row.dividend.percentage)}>{FormatPercentage(row.dividend.percentage)}
                                            </Typography>
                                        </TableCell>
                                        <TableCell key='currencyGain' align="left">
                                            <Typography
                                                variant="body1"
                                                color={getColor(row.currencyGain.amount)}>{FormatCurrency(row.currencyGain.amount, row.currencyGain.currencyCode)}
                                            </Typography>
                                            <Typography
                                                sx={{ opacity: 0.7 }}
                                                variant="body2"
                                                color={getColor(row.currencyGain.percentage)}>{FormatPercentage(row.currencyGain.percentage)}
                                            </Typography>
                                        </TableCell>
                                        <TableCell key='totalReturn' align="left">
                                            <Typography
                                                variant="body1"
                                                color={getColor(row.totalReturn.amount)}>{FormatCurrency(row.totalReturn.amount, row.totalReturn.currencyCode)}
                                            </Typography>
                                            <Typography
                                                sx={{ opacity: 0.7 }}
                                                variant="body2"
                                                color={getColor(row.totalReturn.percentage)}>{FormatPercentage(row.totalReturn.percentage)}
                                            </Typography>
                                        </TableCell>
                                    </TableRow>
                                );
                            })}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                rowsPerPageOptions={[10, 25, 100]}
                component="div"
                count={items.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
            />
        </Paper>
    );
}
