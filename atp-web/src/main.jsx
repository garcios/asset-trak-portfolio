import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import worker from "./mocks/browser";

console.log('Vite mode:', process.env.NODE_ENV);

// Start the worker in mock
if (process.env.NODE_ENV === "mock") {
    console.log("mock environment");
    await worker.start();
}



createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
