const FormatCurrency = (amount: number, currencyCode: string): string => {
    return new Intl.NumberFormat("en-AU", {
        style: "currency",
        currency: currencyCode,
        minimumFractionDigits: 2,
    }).format(amount);
};

export { FormatCurrency };