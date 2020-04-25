export const toRna = (dna) => {
  const rna = {
    G: "C",
    C: "G",
    T: "A",
    A: "U",
  };
  for (i = 0; i < dna.length(); i++) {
    return rna.find(dna[i]).value;
  }
};
