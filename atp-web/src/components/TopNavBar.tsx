import {
    AppBar,
    Toolbar,
    Typography,
    TextField,
    InputAdornment,
    IconButton,
    Box,
    Grid2
} from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import { NavLink } from "react-router-dom";
import React from "react";

// Define navigation links
// Define navigation links
const navLinks = [
    { name: "Holdings", path: "/holdings" },
    { name: "Transactions", path: "/transactions" },
    { name: "Watchlist", path: "/watchlist" },
    { name: "News", path: "/news" },
];

const TopNavBar: React.FC = () => {
    return (
        <Grid2 spacing={2} sx={{ padding: 2 }}>
        <AppBar position="static" color="default" elevation={1}>
            <Toolbar sx={{ display: "flex", justifyContent: "space-between" }}>
                {/* Left: Logo / App Name */}
                <Typography variant="h6" fontWeight="bold" sx={{ flexShrink: 0 }}>
                   Investment Portfolio Navigator
                </Typography>

                {/* Center: Search Bar */}
                <Box sx={{ flexGrow: 1, mx: 4 }}>
                    <TextField
                        variant="outlined"
                        size="small"
                        placeholder="Search stocks..."
                        fullWidth
                        InputProps={{
                            startAdornment: (
                                <InputAdornment position="start">
                                    <SearchIcon />
                                </InputAdornment>
                            ),
                        }}
                        sx={{ backgroundColor: "white", borderRadius: 1 }}
                    />
                </Box>

                {/* Right: Navigation Links & Profile Icon */}
                <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
                    {navLinks.map((link) => (
                        <NavLink
                            key={link.name}
                            to={link.path}
                            style={({ isActive }) => ({
                                textDecoration: "none",
                                color: isActive ? "#1976d2" : "inherit", // Highlight active link in blue
                                fontWeight: isActive ? "bold" : "normal",
                                paddingBottom: isActive ? "2px" : "0",
                                borderBottom: isActive ? "2px solid #1976d2" : "none",
                            })}
                        >
                            {link.name}
                        </NavLink>
                    ))}
                    <IconButton>
                        <AccountCircleIcon fontSize="large" />
                    </IconButton>
                </Box>
            </Toolbar>
        </AppBar>
        </Grid2>
    );
};

export default TopNavBar;
