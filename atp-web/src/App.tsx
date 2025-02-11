import { ApolloProvider } from "@apollo/client";
import Holdings from "./components/Holdings";
import TopNavBar from "./components/TopNavBar";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Transactions from "./components/Transactions";
import Watchlist from "./components/Watchlist";
import News from "./components/News";
import './App.css'
import client from "./services/apollo-client";

function App() {
    return (
        <ApolloProvider client={client}>
            <Router>
                <TopNavBar />
                <Routes>
                    <Route path="/" element={<Holdings />} />
                    <Route path="/holdings" element={<Holdings />} />
                    <Route path="/transactions" element={<Transactions />} />
                    <Route path="/watchlist" element={<Watchlist />} />
                    <Route path="/news" element={<News />} />
                </Routes>
            </Router>
        </ApolloProvider>
    );
}

export default App
