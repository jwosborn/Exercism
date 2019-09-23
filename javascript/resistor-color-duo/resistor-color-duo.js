const bandColors = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white"
];
export const value = ([color1, color2]) => {
  return bandColors.indexOf(color1) * 10 + bandColors.indexOf(color2);
};
