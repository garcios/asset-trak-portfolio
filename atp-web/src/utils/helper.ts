// Helper function to format currency
const FormatCurrency = (amount: number, currencyCode: string): string => {
   return new Intl.NumberFormat("en-AU", {
        style: "currency",
        currency: currencyCode,
        minimumFractionDigits: 2,
    }).format(amount);
};

// Helper function to format percentage
const FormatPercentage = (percent: number) => `${percent.toFixed(2)}%`;

export { FormatCurrency, FormatPercentage };