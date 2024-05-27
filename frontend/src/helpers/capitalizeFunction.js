export default function capitalizeFirstLetterOfEachWord(str) {
  if(str) {
    return str.replace(/\b\w/g, function(char) {
      return char.toUpperCase();
    });
  } else {
    return
  }
}