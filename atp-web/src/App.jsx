import './App.css'
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import BalanceSummary from "./components/BalanceSummary.tsx";
import MasterContainer from "./containers/MasterContainer.tsx";

const client = new ApolloClient({
    uri: "/query",
    cache: new InMemoryCache(),
});

function App() {

    return (
        <ApolloProvider client={client}>
            <MasterContainer/>
        </ApolloProvider>
    );
}

export default App
