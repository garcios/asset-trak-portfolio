import {setupWorker} from "msw/browser";
import handlers from "./handlers";

// Create the worker
const worker = setupWorker(...handlers);

export default worker;