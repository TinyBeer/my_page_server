var eload = document.querySelector('.load_wrapper');
var ebtn = document.querySelector('#btn_add');
var ebox = document.querySelector('.box');
var einp = document.querySelector('#note_content');

window.onload = function () {
    setTimeout(function () {
        console.log("hello");
        eload.classList.add("fade-out");
    }, 3000);
}


ebtn.onclick = function () {
    ebox.insertAdjacentHTML("beforeend",
        `<div class="note">
        <a href="#">
            <span>x</span>
            <p>${einp.value}</p>
        </a>
    </div>`)
    einp.value = "";
    einp.focus();
}
ebox.onclick = function () {
    var tg = event.target;
    if (tg.innerHTML == 'x') {
        tg.parentElement.parentElement.remove();
    }
}