
const concertDates = document.querySelectorAll(".concerts__container li");
console.log(concertDates);
const reversedConcertates = [];
concertDates.forEach((elem) => {
    console.log(elem.innerText);
    const formatedText = elem.innerText.substring(0);
    const array = formatedText.split("-");
    array.reverse();
    elem = array.join("-");
    reversedConcertates.push(elem)
});
console.log(reversedConcertates);
for (let i = 0; i < concertDates.length; i++) {
    const date = new Date(reversedConcertates[i])
    const formatedDate = date.toDateString().substring(4);
    concertDates[i].innerHTML =  formatedDate;
}
// --------------------------------------
function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}
const locationsConcert = document.querySelectorAll(".concerts__container h2");
console.log(locationsConcert);
const formatedLocations = [];
locationsConcert.forEach((elem) => {
    console.log(elem.innerText);
    const array = elem.innerText.split("-");
    const cities = array[0].split("_")
    
    for (let i = 0; i < cities.length; i++) {
        cities[i] = capitalizeFirstLetter(cities[i]);
    }
    cityFormatted = cities.join(" ");
    array[0] = cityFormatted;  
    const countries = array[1].split("_")
    
    for (let i = 0; i < countries.length; i++) {
        countries[i] = capitalizeFirstLetter(countries[i]);
    }
    countryFormatted = countries.join(" ");
    array[1] = countryFormatted;
    elem = array.join(", ");
    formatedLocations.push(elem)
});
console.log(formatedLocations);
for (let i = 0; i < locationsConcert.length; i++) {
    locationsConcert[i].innerHTML =  formatedLocations[i];
}