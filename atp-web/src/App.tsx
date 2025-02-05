import './App.css'
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import HomePage from "./components/HomePage.tsx";

const client = new ApolloClient({
    uri: "/query",
    cache: new InMemoryCache(),
});

function App() {
    return (
        <ApolloProvider client={client}>
            <HomePage/>
        </ApolloProvider>
    );
}

export default App
