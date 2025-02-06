import './App.css'
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import HomePage from "./components/HomePage.tsx";
import TopNavBar from "./components/TopNavBar";
import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Transactions from "./components/Transactions";
import Watchlist from "./components/Watchlist";
import News from "./components/News";

const client = new ApolloClient({
    uri: "/query",
    cache: new InMemoryCache(),
});

function App() {
    return (
        <ApolloProvider client={client}>
            <Router>
                <TopNavBar />
                <Routes>
                    <Route path="/" element={<HomePage />} />
                    <Route path="/holdings" element={<HomePage />} />
                    <Route path="/transactions" element={<Transactions />} />
                    <Route path="/watchlist" element={<Watchlist />} />
                    <Route path="/news" element={<News />} />
                </Routes>
            </Router>
        </ApolloProvider>
    );
}

export default App
