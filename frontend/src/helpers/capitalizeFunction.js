export default function capitalizeFirstLetterOfEachWord(str) {
  return str.replace(/\b\w/g, function(char) {
    return char.toUpperCase();
  });
}