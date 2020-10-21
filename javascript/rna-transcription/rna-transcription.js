const conversion = {
    G: "C",
    C: "G",
    T: "A",
    A: "U",
};

export const toRna = (dnaString) => {
return dnaString.split('').map(function(dna) {
        return conversion[dna];
    }).join('');
};
