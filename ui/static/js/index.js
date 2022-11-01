const releaseDates = document.querySelectorAll(".aritist-album-realese b");
console.log(releaseDates);
const reversedReleaseDates = [];
releaseDates.forEach((elem) => {
    console.log(elem.innerText);
    const array = elem.innerText.split("-");
    array.reverse();
    elem = array.join("-");
    reversedReleaseDates.push(elem)
});

console.log(reversedReleaseDates);

for (let i = 0; i < releaseDates.length; i++) {
    const date = new Date(reversedReleaseDates[i])
    const formatedDate = date.toDateString().substring(4);
    releaseDates[i].innerHTML =  formatedDate;
}