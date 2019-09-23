const GS = 1000000000;
export const gigasecond = date => {
  return new Date(date.getTime() + GS * 1000);
};
